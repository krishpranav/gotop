package utils

import (
	"sort"
	"strconv"
)

func SortData(data [][]string, sortIdx int, sortAsc bool, sortCase string) {

	intSort := func(i, j int) bool {
		x, _ := strconv.Atoi(data[i][sortIdx])
		y, _ := strconv.Atoi(data[j][sortIdx])
		if sortAsc {
			return x < y
		}
		return x > y
	}

	strSort := func(i, j int) bool {
		if sortAsc {
			return data[i][sortIdx] < data[j][sortIdx]
		}
		return data[i][sortIdx] > data[j][sortIdx]
	}

	floatSort := func(i, j int) bool {
		x1 := data[i][sortIdx]
		y1 := data[j][sortIdx]
		x, _ := strconv.ParseFloat(x1[:len(x1)-1], 32)
		y, _ := strconv.ParseFloat(y1[:len(y1)-1], 32)
		if sortAsc {
			return x < y
		}
		return x > y
	}

	sortFuncs := make(map[int]func(i, j int) bool)
	switch sortCase {
	case "PROCS":
		sortFuncs = map[int]func(i, j int) bool{
			0: intSort,
			1: strSort,
			3: floatSort,
			2: floatSort,
			4: strSort,
			5: strSort,
			6: strSort,
			7: intSort,
		}
	case "CONTAINER":
		sortFuncs = map[int]func(i, j int) bool{
			0: strSort,
			1: strSort,
			2: strSort,
			3: strSort,
			4: strSort,
			5: floatSort,
			6: floatSort,
		}

	default:
		sortFuncs[sortIdx] = strSort
	}

	sort.Slice(data, sortFuncs[sortIdx])
}
