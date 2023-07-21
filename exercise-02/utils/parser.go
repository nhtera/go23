package utils

import (
	"strconv"
)

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
