package utils

import "strconv"

func MustParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}
