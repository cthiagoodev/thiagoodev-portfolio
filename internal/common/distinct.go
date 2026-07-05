package common

func Distinct[T comparable](slice []T) []T {
	viewed := make(map[T]bool, len(slice))
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if !viewed[v] {
			viewed[v] = true
			result = append(result, v)
		}
	}

	return result
}
