package iteration

func Repeat(character string, repetition int) (repeated string) {
	for i := 0; i < repetition; i++ {
		repeated = repeated + character
	}
	return repeated
}