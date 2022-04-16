package test

import (
	leighmcculloch "4d63.com/optional"
	"testing"
)

func Benchmark_Get_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = leighmcculloch.Of("some string").Get()
	}
}

func Benchmark_Get_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = leighmcculloch.Empty[string]().Get()
	}
}

func Benchmark_IsPresent_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Of("some string").IsPresent()
	}
}

func Benchmark_IsPresent_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Empty[string]().IsPresent()
	}
}

func Benchmark_IfPresent_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		leighmcculloch.Of("some string").If(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		leighmcculloch.Empty[string]().If(func(s string) {})
	}
}

func Benchmark_Else_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Of("some string").Else("other string")
	}
}

func Benchmark_Else_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Empty[string]().Else("other string")
	}
}

func Benchmark_ElseGet_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Of("some string").ElseFunc(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Empty[string]().ElseFunc(func() string { return "other string" })
	}
}

func Benchmark_ElseZero_Present_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Of("some string").ElseZero()
	}
}

func Benchmark_ElseZero_Empty_leighmcculloch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = leighmcculloch.Empty[string]().ElseZero()
	}
}
