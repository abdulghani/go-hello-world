package utils

func ShouldPanic(err error) {
	if err != nil {
		panic(err)
	}
}
