package utils

func MapToSlice[S comparable, T any](m map[S]T) []T {
	slice := make([]T, 0, len(m))
	for _, t := range m {
		slice = append(slice, t)
	}
	return slice
}
