package testunitario

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//go test -cover
//go test -coverprofile=coverage
//go tool cover -html=coverage
//go tool cover -func=coverage