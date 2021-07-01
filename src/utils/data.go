package utils

type DataStats struct {
	NetStats  map[string][]float64
	FieldSet  string
	CpuStats  []float64
	MemStats  []float64
	DiskStats [][]string
}
