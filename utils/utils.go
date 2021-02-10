package utils

// CheckError panics error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}