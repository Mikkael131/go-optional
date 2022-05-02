package optional

import (
	"testing"
)

func Benchmark_Init_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string")
	}
}

func Benchmark_Init_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]
	}
}

func Benchmark_Init_OfPtr_Value(b *testing.B) {
	s := "some string"
	for n := 0; n < b.N; n++ {
		_ = OfPtr(&s)
	}
}

func Benchmark_Init_OfPtr_Nil(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = OfPtr(new(string))
	}
}
