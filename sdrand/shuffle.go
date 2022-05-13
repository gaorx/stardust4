package sdrand

import (
	"math/rand"
)

func Shuffle[T any](l []T) {
	n := len(l)
	if n <= 1 {
		return
	}
	clone := make([]T, n)
	for i := 0; i < n; i++ {
		clone[i] = l[i]
	}
	perms := rand.Perm(n)
	for i := 0; i < n; i++ {
		l[i] = clone[perms[i]]
	}
}
