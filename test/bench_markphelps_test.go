package test

import (
	markphelps "github.com/markphelps/optional"
	"testing"
)

func Benchmark_Get_Present_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = markphelps.NewString("some string").Get()
	}
}

func Benchmark_Get_Empty_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = markphelps.String{}.Get()
	}
}

func Benchmark_IsPresent_Present_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.NewString("some string").Present()
	}
}

func Benchmark_IsPresent_Empty_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.String{}.Present()
	}
}

func Benchmark_IfPresent_Present_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		markphelps.NewString("some string").If(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		markphelps.String{}.If(func(s string) {})
	}
}

func Benchmark_Else_Present_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.NewString("some string").OrElse("other string")
	}
}

func Benchmark_Else_Empty_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.String{}.OrElse("other string")
	}
}
