package rslicer

import "testing"

var exampleSmall string = "😸😼"                                          // len = 2
var exampleMedium string = "😸 Some very long string with nice unicode😼" // len = 42
// len = 413
var exampleLarge string = "😸 Чёрная дыра́ — область пространства-времени, гравитационное притяжение которой настолько велико, что покинуть её не могут даже объекты, движущиеся со скоростью света, в том числе кванты самого света. Граница этой области называется горизонтом событий. В простейшем случае сферически симметричной чёрной дыры он представляет собой сферу с радиусом Шварцшильда, который считается характерным размером чёрной дыры."

const (
	smallBegin     int = 0
	smallEnd       int = 1
	smallNegBegin  int = -2
	smallNegEnd    int = -1
	mediumBegin    int = 4
	mediumEnd      int = 18
	mediumNegBegin int = -38
	mediumNegEnd   int = -20
	largeBegin     int = 10
	largeEnd       int = 120
	largeNegBegin  int = -400
	largeNegEnd    int = -380
	largeEqual     int = 410
)

func BenchmarkPositiveSmallWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleSmall, smallBegin, smallEnd)
	}
}
func BenchmarkNegativeSmallWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleSmall, smallNegBegin, smallNegEnd)
	}
}

func BenchmarkPositiveMediumWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleMedium, mediumBegin, mediumEnd)
	}
}

func BenchmarkNegativeMediumWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleMedium, mediumNegBegin, mediumNegEnd)
	}
}

func BenchmarkPositiveLargeWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleLarge, largeBegin, largeEnd)
	}
}

func BenchmarkNegativeLargeWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = GetRuneRange(exampleLarge, largeNegBegin, largeNegEnd)
	}
}

func BenchmarkNegativeLargeSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetRuneSlice(exampleLarge, largeNegBegin, largeNegEnd)
	}
}
