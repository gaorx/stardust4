package sdload

func MustBytes(loc string) []byte {
	data, err := Bytes(loc)
	if err != nil {
		panic(err)
	}
	return data
}

func MustText(loc string) string {
	s, err := Text(loc)
	if err != nil {
		panic(err)
	}
	return s
}
