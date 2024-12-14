package utils

func PointerValue[T any](val T) *T {
	return &val
}
