package container

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/krishpranav/gotop/src/utils"
)

func TestGetPerCPUPercents(t *testing.T) {
	tests := []struct {
		inputStats     types.StatsJSON
		expectedOutput []string
	}{
		{
			inputStats: types.StatsJSON{
				Stats: types.Stats{
					CPUStats: types.CPUStats{
						CPUUsage: types.CPUUsage{
							PercpuUsage: []uint64{0},
						},
					},
					PreCPUStats: types.CPUStats{
						CPUUsage: types.CPUUsage{
							PercpuUsage: []uint64{0, 0},
						},
					},
				},
			},
			expectedOutput: []string{"NA", "NA"},
		},
	}

	for _, test := range tests {
		testVal := getPerCPUPercents(&test.inputStats)
		utils.Equals(t, testVal, test.expectedOutput)
	}
}