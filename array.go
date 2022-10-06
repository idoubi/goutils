package goutils

// IsStrInArr: find element in array with string type
func IsStrInArr(arr []string, str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}

	return false
}

// IsIntInArr: find element in array with int type
func IsIntInArr(arr []int, i int) bool {
	for _, v := range arr {
		if i == v {
			return true
		}
	}

	return false
}

// IsInt64InArr: find element in array with int64 type
func IsInt64InArr(arr []int64, i int64) bool {
	for _, v := range arr {
		if i == v {
			return true
		}
	}

	return false
}
