package sdlo

func SliceNilAsEmpty[T any](s []T) []T {
	if s == nil {
		return make([]T, 0)
	}
	return s
}

func SliceCopy[T any](src []T) []T {
	if src == nil {
		return nil
	}
	n := len(src)
	dst := make([]T, n)
	copy(dst, src)
	return dst
}

func SliceEqual[T comparable](a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	aLen, bLen := len(a), len(b)
	if aLen != bLen {
		return false
	}
	for i := 0; i < aLen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
