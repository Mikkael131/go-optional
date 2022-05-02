package test

import (
	leighmcculloch "4d63.com/optional"
	"testing"
)

func Benchmark_Init_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Of("some string")
	}
}

func Benchmark_Init_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Empty[string]
	}
}

func Benchmark_Init_OfPtr_Value_leighmcculloch(b *testing.B) {
	s := "some string"
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.OfPtr(&s)
	}
}

func Benchmark_Init_OfPtr_Nil_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.OfPtr(new(string))
	}
}

func Benchmark_Get_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_Get_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_IsPresent_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsPresent()
	}
}

func Benchmark_IsPresent_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsPresent()
	}
}

func Benchmark_IfPresent_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.If(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.If(func(s string) {})
	}
}

func Benchmark_Else_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Else("other string")
	}
}

func Benchmark_Else_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Else("other string")
	}
}

func Benchmark_ElseGet_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseFunc(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseFunc(func() string { return "other string" })
	}
}

func Benchmark_ElseZero_Present_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseZero()
	}
}

func Benchmark_ElseZero_Empty_leighmcculloch(b *testing.B) {
	b.StopTimer()
	o := leighmcculloch.Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseZero()
	}
}
