package test

import (
	moznion "github.com/moznion/go-optional"
	"testing"
)

func Benchmark_Get_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = moznion.Some("some string").Take()
	}
}

func Benchmark_Get_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = moznion.None[string]().Take()
	}
}

func Benchmark_IsPresent_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.Some("some string").IsSome()
	}
}

func Benchmark_IsPresent_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.None[string]().IsSome()
	}
}

func Benchmark_Else_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.Some("some string").TakeOr("other string")
	}
}

func Benchmark_Else_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.None[string]().TakeOr("other string")
	}
}

func Benchmark_ElseGet_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.Some("some string").TakeOrElse(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.None[string]().TakeOrElse(func() string { return "other string" })
	}
}

func Benchmark_Filter_Present_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.Some("some string").Filter(func(s string) bool {
			return true
		})
		_ = moznion.Some("some string").Filter(func(s string) bool {
			return false
		})
	}
}

func Benchmark_Filter_Empty_moznion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = moznion.None[string]().Filter(func(s string) bool {
			return true
		})
		_ = moznion.None[string]().Filter(func(s string) bool {
			return false
		})
	}
}
