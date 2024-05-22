package rslicer

import "testing"

func BenchmarkSmallThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleSmall)
		part := runes[smallBegin:smallEnd]
		_ = string(part)
	}
}

func BenchmarkMediumThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleMedium)
		part := runes[mediumBegin:mediumEnd]
		_ = string(part)
	}
}

func BenchmarkLargeThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleLarge)
		part := runes[largeBegin:largeEnd]
		_ = string(part)
	}
}
