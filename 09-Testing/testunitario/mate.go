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

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}


//go test -cpuprofile=cou
//go tool pprof cou
//top
//list Fibonacci
//quit