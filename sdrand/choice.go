package sdrand

import "math/rand"

type W[T any] struct {
	W int `json:"w"`
	V T   `json:"v"`
}

func ChoiceOne[T any](choices ...T) T {
	var def T
	n := len(choices)
	if n <= 0 {
		return def
	}
	return choices[rand.Intn(n)]
}

func ChoiceOneW[T any](choices ...W[T]) T {
	var def T
	n := len(choices)
	if n <= 0 {
		return def
	}
	if n == 1 {
		first := choices[0]
		if first.W > 0 {
			return first.V
		} else {
			return def
		}
	}
	var sum, upto int64 = 0, 0
	for _, w := range choices {
		if w.W > 0 {
			sum += int64(w.W)
		}
	}
	r := Float64r(0.0, float64(sum))
	for _, w := range choices {
		ww := w.W
		if ww < 0 {
			ww = 0
		}
		if float64(upto)+float64(ww) > r {
			return w.V
		}
		upto += int64(w.W)
	}
	return def
}

func ChoiceSome[T any](choices []T, n int) []T {
	nChoice := len(choices)
	if nChoice == 0 || n <= 0 {
		return []T{}
	}
	m := make(map[int]T, nChoice)
	for i, v := range choices {
		m[i] = v
	}
	r := make([]T, 0, n)
	for i := 0; i < n; i++ {
		if len(m) <= 0 || len(r) >= n {
			break
		}
		c := rand.Intn(len(m))
		j := 0
	Next:
		for arrIndex, v := range m {
			if j == c {
				r = append(r, v)
				delete(m, arrIndex)
				break Next
			}
			j++
		}
	}
	return r
}
