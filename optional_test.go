package optional

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptional_Get(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     string
		wantOk   bool
	}{
		"get present value": {
			optional: Of("some string"),
			want:     "some string",
			wantOk:   true,
		},
		"get non present value": {
			optional: Empty[string](),
			want:     "",
			wantOk:   false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotOk := tt.optional.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestOptional_IsPresent(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     bool
	}{
		"false with non present value": {
			optional: Empty[string](),
			want:     false,
		},
		"true with present value": {
			optional: Of("some string"),
			want:     true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.IsPresent(); got != tt.want {
				t.Errorf("IsPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_IfPresent(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		wantCalled bool
	}{
		"not called with non present value": {
			optional:   Empty[string](),
			wantCalled: false,
		},
		"called with present value": {
			optional:   Of("some string"),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			tt.optional.IfPresent(func(v string) {
				called = true
			})
			if called != tt.wantCalled {
				t.Errorf("IsPresent() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}

func TestOptional_Else(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		arg      string
		want     string
	}{
		"value of argument with empty optional": {
			optional: Empty[string](),
			arg:      "some string",
			want:     "some string",
		},
		"value of present optional": {
			optional: Of("some string"),
			arg:      "other string",
			want:     "some string",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.Else(tt.arg); got != tt.want {
				t.Errorf("Else() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_ElseGet(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		arg      func() string
		want     string
	}{
		"value of function with empty optional": {
			optional: Empty[string](),
			arg:      func() string { return "some string" },
			want:     "some string",
		},
		"value of present optional": {
			optional: Of("some string"),
			arg:      func() string { return "other string" },
			want:     "some string",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.ElseGet(tt.arg); got != tt.want {
				t.Errorf("ElseGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_ElseErr(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		arg      error
		want     string
		wantErr  error
	}{
		"nil error from of argument with empty optional": {
			optional: Empty[string](),
			arg:      nil,
			want:     "",
			wantErr:  nil,
		},
		"error from of argument with empty optional": {
			optional: Empty[string](),
			arg:      fmt.Errorf("some error"),
			want:     "",
			wantErr:  fmt.Errorf("some error"),
		},
		"value of present optional": {
			optional: Of("some string"),
			arg:      fmt.Errorf("some error"),
			want:     "some string",
			wantErr:  nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := tt.optional.ElseErr(tt.arg)
			if got != tt.want {
				t.Errorf("ElseErr() got = %v, want %v", got, tt.want)
			}
			if (gotErr == nil && tt.wantErr != nil) ||
				(gotErr != nil && tt.wantErr == nil) ||
				(gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("ElseErr() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestOptional_Filter(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		arg        func(v string) bool
		want       Optional[string]
		wantCalled bool
	}{
		"do not filter empty optional": {
			optional:   Empty[string](),
			arg:        nil,
			want:       Empty[string](),
			wantCalled: false,
		},
		"filter present optional with matching predicate": {
			optional: Of("some string"),
			arg: func(v string) bool {
				return v == "some string"
			},
			want:       Of("some string"),
			wantCalled: true,
		},
		"filter present optional with non-matching predicate": {
			optional: Of("some string"),
			arg: func(v string) bool {
				return v == "other string"
			},
			want:       Empty[string](),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			argWrapper := func(v string) bool {
				called = true
				return tt.arg(v)
			}

			got := tt.optional.Filter(argWrapper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
			if called != tt.wantCalled {
				t.Errorf("Filter() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}

func TestOptional_Map(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		arg        func(v string) (string, bool)
		want       Optional[string]
		wantCalled bool
	}{
		"do not map empty optional": {
			optional:   Empty[string](),
			arg:        nil,
			want:       Empty[string](),
			wantCalled: false,
		},
		"map present optional with returned ok": {
			optional: Of("some string"),
			arg: func(v string) (string, bool) {
				return "other string", true
			},
			want:       Of("other string"),
			wantCalled: true,
		},
		"do not map present optional with returned nok": {
			optional: Of("some string"),
			arg: func(v string) (string, bool) {
				return "other string", false
			},
			want:       Empty[string](),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			argWrapper := func(v string) (string, bool) {
				called = true
				return tt.arg(v)
			}

			got := tt.optional.Map(argWrapper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
			if called != tt.wantCalled {
				t.Errorf("Map() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}

func TestOptional_FlatMap(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		arg        func(v string) Optional[string]
		want       Optional[string]
		wantCalled bool
	}{
		"do not flat map empty optional": {
			optional:   Empty[string](),
			arg:        nil,
			want:       Empty[string](),
			wantCalled: false,
		},
		"flat map present optional": {
			optional: Of("some string"),
			arg: func(v string) Optional[string] {
				return Of("other string")
			},
			want:       Of("other string"),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			argWrapper := func(v string) Optional[string] {
				called = true
				return tt.arg(v)
			}

			got := tt.optional.FlatMap(argWrapper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
			if called != tt.wantCalled {
				t.Errorf("FlatMap() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}

func TestOptional_String(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     string
	}{
		"empty optional": {
			optional: Empty[string](),
			want:     "Optional[empty]",
		},
		"present optional": {
			optional: Of("some string"),
			want:     "Optional[some string]",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
