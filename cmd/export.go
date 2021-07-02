package cmd

import (
	"fmt"
	"strings"

	exportGeneral "github.com/krishpranav/gotop/src/export/general"
	exportProc "github.com/krishpranav/gotop/src/export/proc"

	"github.com/spf13/cobra"
)

const (
	defaultExportRefreshRate = 1000
	defaultExportIterations  = 10
	defaultExportFileName    = "gotop_profile"
	defaultExportType        = "json"
	defaultExportPid         = -1
)

var providedExportTypes = map[string]bool{
	"json": true,
}

func hasValidExtension(filename, exportType string) error {
	filename = strings.ToLower(filename)

	var hasProvidedExtension bool = false

	for exportType, allowed := range providedExportTypes {
		if allowed {
			hasType := strings.HasSuffix(filename, "."+exportType)
			hasProvidedExtension = hasProvidedExtension || hasType
		}
	}

	if hasProvidedExtension {
		validExtension := strings.HasSuffix(filename, exportType)
		if validExtension {
			return nil
		}
		return fmt.Errorf("invaid file extension")
	}

	return nil
}

func validateFileName(filename, exportType string) error {
	isValid := hasValidExtension(filename, exportType)
	return isValid
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Used to export profiled data.",
	Long:  `the export command can be used to export profiled data to a specific file format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		iter, err := cmd.Flags().GetUint32("iter")
		if err != nil {
			return err
		}

		refreshRate, err := cmd.Flags().GetUint64("refresh")
		if err != nil {
			return err
		}

		exportType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		exportType = strings.ToLower(exportType)
		if validExportType := providedExportTypes[exportType]; !validExportType {
			return fmt.Errorf("export type not supported")
		}

		exportPid, err := cmd.Flags().GetInt32("pid")
		if err != nil {
			return err
		}

		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}

		err = validateFileName(filename, exportType)
		if err != nil {
			return err
		}

		if exportPid == defaultExportPid {
			switch exportType {
			case "json":
				return exportGeneral.ExportJSON(filename, iter, refreshRate)

			default:
				return fmt.Errorf("invalid export type, see gotop export --help")
			}
		} else {
			switch exportType {
			case "json":
				return exportProc.ExportPidJSON(exportPid, filename, iter, refreshRate)

			default:
				return fmt.Errorf("invalid export type, see gotop export --help")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().Uint32P(
		"iter",
		"i",
		defaultExportIterations,
		"specify the number of iterations to run profiler",
	)
	exportCmd.Flags().StringP(
		"filename",
		"f",
		defaultExportFileName,
		"specify the name of the export file",
	)
	exportCmd.Flags().Uint64P(
		"refresh",
		"r",
		defaultExportRefreshRate,
		"specify frequency of data fetch in milliseconds",
	)
	exportCmd.Flags().StringP(
		"type",
		"t",
		defaultExportType,
		"specify the output format of the profiling result (json by default)",
	)
	exportCmd.Flags().Int32P(
		"pid",
		"p",
		defaultExportPid,
		"specify pid of process to profile, ignore to profile all processes",
	)
}
