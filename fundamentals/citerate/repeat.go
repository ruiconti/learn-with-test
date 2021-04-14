package iteration

const repeatCount = 5

func Repeat(character string, N int) (repeated string) {
	if N == 0 {
		N = repeatCount
	}
	for i := 0; i < N; i++ {
		repeated += character
	}
	return
}
