package optional

import (
	"fmt"
	"testing"
)

func Benchmark_Get_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Of("some string").Get()
	}
}

func Benchmark_Get_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Empty[string]().Get()
	}
}

func Benchmark_IsPresent_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").IsPresent()
	}
}

func Benchmark_IsPresent_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().IsPresent()
	}
}

func Benchmark_IfPresent_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Of("some string").IfPresent(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Empty[string]().IfPresent(func(s string) {})
	}
}

func Benchmark_Else_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").Else("other string")
	}
}

func Benchmark_Else_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().Else("other string")
	}
}

func Benchmark_ElseGet_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").ElseGet(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().ElseGet(func() string { return "other string" })
	}
}

func Benchmark_ElseErr_Present(b *testing.B) {
	err := fmt.Errorf("some error")
	for n := 0; n < b.N; n++ {
		_, _ = Of("some string").ElseErr(err)
	}
}

func Benchmark_ElseErr_Empty(b *testing.B) {
	err := fmt.Errorf("some error")
	for n := 0; n < b.N; n++ {
		_, _ = Empty[string]().ElseErr(err)
	}
}

func Benchmark_ElseZero_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").ElseZero()
	}
}

func Benchmark_ElseZero_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().ElseZero()
	}
}

func Benchmark_Filter_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").Filter(func(s string) bool {
			return true
		})
		_ = Of("some string").Filter(func(s string) bool {
			return false
		})
	}
}

func Benchmark_Filter_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().Filter(func(s string) bool {
			return true
		})
		_ = Empty[string]().Filter(func(s string) bool {
			return false
		})
	}
}

func Benchmark_Map_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").Map(func(s string) (string, bool) {
			return "other string", true
		})
		_ = Of("some string").Map(func(s string) (string, bool) {
			return "", false
		})
	}
}

func Benchmark_Map_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().Map(func(s string) (string, bool) {
			return "other string", true
		})
		_ = Empty[string]().Map(func(s string) (string, bool) {
			return "", false
		})
	}
}

func Benchmark_FlatMap_Present(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Of("some string").FlatMap(func(s string) Optional[string] {
			return Of("other string")
		})
		_ = Of("some string").FlatMap(func(s string) Optional[string] {
			return Empty[string]()
		})
	}
}

func Benchmark_FlatMap_Empty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Empty[string]().FlatMap(func(s string) Optional[string] {
			return Of("other string")
		})
		_ = Empty[string]().FlatMap(func(s string) Optional[string] {
			return Empty[string]()
		})
	}
}
