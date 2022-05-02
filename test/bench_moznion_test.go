package test

import (
	moznion "github.com/moznion/go-optional"
	"testing"
)

func Benchmark_Init_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.Some("some string")
	}
}

func Benchmark_Init_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.None[string]()
	}
}

func Benchmark_Get_Present_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.Some("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Take()
	}
}

func Benchmark_Get_Empty_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.None[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Take()
	}
}

func Benchmark_IsPresent_Present_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.Some("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsSome()
	}
}

func Benchmark_IsPresent_Empty_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.None[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsSome()
	}
}

func Benchmark_Else_Present_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.Some("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.TakeOr("other string")
	}
}

func Benchmark_Else_Empty_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.None[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.TakeOr("other string")
	}
}

func Benchmark_ElseGet_Present_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.Some("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.TakeOrElse(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.None[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.TakeOrElse(func() string { return "other string" })
	}
}

func Benchmark_Filter_Present_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.Some("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Filter(func(s string) bool {
			return true
		})
		_ = o.Filter(func(s string) bool {
			return false
		})
	}
}

func Benchmark_Filter_Empty_moznion(b *testing.B) {
	b.StopTimer()
	o := moznion.None[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Filter(func(s string) bool {
			return true
		})
		_ = o.Filter(func(s string) bool {
			return false
		})
	}
}
