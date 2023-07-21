package utils

import (
	"fmt"
	"sort"
)

func SortInts(ints []int) {
	sort.Ints(ints)
}

func SortFloats(floats []float64) {
	sort.Float64s(floats)
}

func SortStrings(strings []string) {
	sort.Strings(strings)
}

func SortMixed(mixed []interface{}) {
	sort.Slice(mixed, func(i, j int) bool {
		switch a := mixed[i].(type) {
		case int:
			b := mixed[j].(int)
			return a < b
		case float64:
			b := mixed[j].(float64)
			return a < b
		case string:
			b := mixed[j].(string)
			return a < b
		}
		return false
	})
}

func Sort(inputType string, input []string) []string {
	switch inputType {
	case "int":
		ints := make([]int, len(input))
		for i, v := range input {
			ints[i] = Atoi(v)
		}
		SortInts(ints)
		result := make([]string, len(ints))
		for i, v := range ints {
			result[i] = fmt.Sprintf("%d", v)
		}
		return result
	case "float":
		floats := make([]float64, len(input))
		for i, v := range input {
			floats[i] = Atof(v)
		}
		SortFloats(floats)
		result := make([]string, len(floats))
		for i, v := range floats {
			result[i] = fmt.Sprintf("%g", v)
		}
		return result
	case "string":
		SortStrings(input)
		return input
	case "mix":
		mixed := make([]interface{}, len(input))
		for i, v := range input {
			if IsInt(v) {
				// mixed[i] = Atoi(v)
				mixed[i] = v
			} else if IsFloat(v) {
				// mixed[i] = Atof(v)
				mixed[i] = v
			} else {
				mixed[i] = v
			}
		}
		SortMixed(mixed)
		result := make([]string, len(mixed))
		for i, v := range mixed {
			result[i] = fmt.Sprintf("%v", v)
		}
		return result
	}
	return nil
}
