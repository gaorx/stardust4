package sdlo

func IfEmptyElse[T comparable](v, def T) T {
	var empty T
	if v == empty {
		return def
	} else {
		return v
	}
}
