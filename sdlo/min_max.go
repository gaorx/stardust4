package sdlo

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func Max2[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min2[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxN[T constraints.Ordered](values ...T) T {
	return lo.Max[T](values)
}

func MinN[T constraints.Ordered](values ...T) T {
	return lo.Min[T](values)
}
