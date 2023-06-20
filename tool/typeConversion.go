package tool

import "strconv"

//------------- float -------------

func Float64ToInt64(in float64) int64 {
	return int64(in)
}

func Float64ToString(in float64, prec int) string {
	return strconv.FormatFloat(in, 'f', prec, 64)
}

func StringToFloat64(in string) float64 {
	out, _ := strconv.ParseFloat(in, 64)
	return out
}