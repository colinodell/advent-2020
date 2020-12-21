package utils

import "reflect"

func SliceContains(haystack interface{}, needle interface{}) bool {
	h := reflect.ValueOf(haystack)

	for i := 0; i < h.Len(); i++ {
		if h.Index(i).Interface() == needle {
			return true
		}
	}

	return false
}

// Based on https://github.com/juliangruber/go-intersect/blob/master/intersect.go
func IntersectStrings(a []string, b []string) []string {
	set := make([]string, 0)

	for _, v := range a {
		if SliceContains(b, v) {
			set = append(set, v)
		}
	}

	return set
}


func RemoveItem(haystack []string, needle string) []string {
	var newitems []string

	for _, i := range haystack {
		if i != needle {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
