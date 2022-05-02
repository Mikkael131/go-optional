package test

import (
	markphelps "github.com/markphelps/optional"
	"testing"
)

func Benchmark_Init_Present_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.NewString("some string")
	}
}

func Benchmark_Init_Empty_markphelps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = markphelps.String{}
	}
}

func Benchmark_Get_Present_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.NewString("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_Get_Empty_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.String{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_IsPresent_Present_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.NewString("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Present()
	}
}

func Benchmark_IsPresent_Empty_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.String{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Present()
	}
}

func Benchmark_IfPresent_Present_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.NewString("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.If(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.String{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.If(func(s string) {})
	}
}

func Benchmark_Else_Present_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.NewString("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.OrElse("other string")
	}
}

func Benchmark_Else_Empty_markphelps(b *testing.B) {
	b.StopTimer()
	o := markphelps.String{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.OrElse("other string")
	}
}
