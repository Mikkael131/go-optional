package optional

import (
	"fmt"
	"testing"
)

func Benchmark_Get_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_Get_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.Get()
	}
}

func Benchmark_IsPresent_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsPresent()
	}
}

func Benchmark_IsPresent_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.IsPresent()
	}
}

func Benchmark_IfPresent_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.IfPresent(func(s string) {})
	}
}

func Benchmark_IfPresent_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.IfPresent(func(s string) {})
	}
}

func Benchmark_IfPresentOrElse_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.IfPresentOrElse(func(s string) {}, func() {})
	}
}

func Benchmark_IfPresentOrElse_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		o.IfPresentOrElse(func(s string) {}, func() {})
	}
}

func Benchmark_Else_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Else("other string")
	}
}

func Benchmark_Else_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Else("other string")
	}
}

func Benchmark_ElseGet_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseGet(func() string { return "other string" })
	}
}

func Benchmark_ElseGet_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseGet(func() string { return "other string" })
	}
}

func Benchmark_ElseErr_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	err := fmt.Errorf("some error")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.ElseErr(err)
	}
}

func Benchmark_ElseErr_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	err := fmt.Errorf("some error")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.ElseErr(err)
	}
}

func Benchmark_ElseZero_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseZero()
	}
}

func Benchmark_ElseZero_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.ElseZero()
	}
}

func Benchmark_Filter_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
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

func Benchmark_Filter_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
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

func Benchmark_Map_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Map(func(s string) (string, bool) {
			return "other string", true
		})
		_ = o.Map(func(s string) (string, bool) {
			return "", false
		})
	}
}

func Benchmark_Map_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Map(func(s string) (string, bool) {
			return "other string", true
		})
		_ = o.Map(func(s string) (string, bool) {
			return "", false
		})
	}
}

func Benchmark_FlatMap_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.FlatMap(func(s string) Optional[string] {
			return Of("other string")
		})
		_ = o.FlatMap(func(s string) Optional[string] {
			return Empty[string]()
		})
	}
}

func Benchmark_FlatMap_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.FlatMap(func(s string) Optional[string] {
			return Of("other string")
		})
		_ = o.FlatMap(func(s string) Optional[string] {
			return Empty[string]()
		})
	}
}

func Benchmark_Ptr_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Ptr()
	}
}

func Benchmark_Ptr_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Ptr()
	}
}

func Benchmark_Val_Present(b *testing.B) {
	b.StopTimer()
	o := ptr(Of("some string"))
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Val()
	}
}

func Benchmark_Val_Empty(b *testing.B) {
	b.StopTimer()
	o := ptr(Empty[string]())
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.Val()
	}
}

func Benchmark_MarshalJSON_Present(b *testing.B) {
	b.StopTimer()
	o := Of("some string")
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.MarshalJSON()
	}
}

func Benchmark_MarshalJSON_Empty(b *testing.B) {
	b.StopTimer()
	o := Empty[string]()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = o.MarshalJSON()
	}
}

func Benchmark_UnmarshalJSON_String(b *testing.B) {
	b.StopTimer()
	bytes := []byte("some string")
	o := &Optional[string]{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.UnmarshalJSON(bytes)
	}
}

func Benchmark_UnmarshalJSON_EmptyString(b *testing.B) {
	b.StopTimer()
	bytes := []byte("")
	o := &Optional[string]{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.UnmarshalJSON(bytes)
	}
}

func Benchmark_UnmarshalJSON_NullString(b *testing.B) {
	b.StopTimer()
	bytes := []byte("null")
	o := &Optional[string]{}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = o.UnmarshalJSON(bytes)
	}
}
