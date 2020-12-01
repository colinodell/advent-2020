package utils

func SliceContains(haystack []int, needle int) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}

	return false
}
