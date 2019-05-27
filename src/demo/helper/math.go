package helper

func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	// v |= v >> 4
	// v |= v >> 8
	// v |= v >> 16
	v++
	return v
}

func fun(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fun(n-1) + fun(n-2)
}
