package utils

// SearchStringInArray func for searching value in []string array.
func SearchStringInArray(value string, array []string) (ok bool) {
	for index := range array {
		if ok = array[index] == value; ok {
			return
		}
	}

	return
}
