package tool

import "strconv"

//------------- float -------------

func Float64ToInt64(in float64) int64 {
	return int64(in)
}

func Float64ToString(in float64, prec int) string {
	return strconv.FormatFloat(in, 'f', prec, 64)
}