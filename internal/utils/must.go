package utils

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
