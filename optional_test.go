package optional

import (
	"encoding/json"
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

func TestOptional_ElseZero(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     string
	}{
		"zero of empty optional": {
			optional: Empty[string](),
			want:     "",
		},
		"value of present optional": {
			optional: Of("some string"),
			want:     "some string",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.optional.ElseZero()
			if got != tt.want {
				t.Errorf("ElseZero() got = %v, want %v", got, tt.want)
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

func TestOptional_Ptr(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     *Optional[string]
	}{
		"empty optional": {
			optional: Empty[string](),
			want:     nil,
		},
		"present optional": {
			optional: Of("some string"),
			want:     ptr(Of("some string")),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.Ptr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_Val(t *testing.T) {
	tests := map[string]struct {
		optional *Optional[string]
		want     Optional[string]
	}{
		"nil": {
			optional: nil,
			want:     Empty[string](),
		},
		"empty optional": {
			optional: ptr(Empty[string]()),
			want:     Empty[string](),
		},
		"present optional": {
			optional: ptr(Of("some string")),
			want:     Of("some string"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.optional.Val(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Val() = %v, want %v", got, tt.want)
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

func TestOptional_MarshalJSON_Func(t *testing.T) {
	tests := map[string]struct {
		optional Optional[string]
		want     []byte
		wantErr  bool
	}{
		"null with empty optional": {
			optional: Empty[string](),
			want:     []byte("null"),
		},
		"value with present optional": {
			optional: Of("some string"),
			want:     []byte("\"some string\""),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.optional.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_MarshalJSON_StructWithValue(t *testing.T) {
	tests := map[string]struct {
		arg     interface{}
		want    string
		wantErr bool
	}{
		"null with uninitialized optional string": {
			arg:  valueStruct[string]{},
			want: `{"v":null}`,
		},
		"null with uninitialized optional int": {
			arg:  valueStruct[int]{},
			want: `{"v":null}`,
		},
		"null with uninitialized optional float64": {
			arg:  valueStruct[float64]{},
			want: `{"v":null}`,
		},
		"null with empty optional string": {
			arg:  valueStruct[string]{V: Empty[string]()},
			want: `{"v":null}`,
		},
		"null with empty optional int": {
			arg:  valueStruct[int]{V: Empty[int]()},
			want: `{"v":null}`,
		},
		"null with empty optional float64": {
			arg:  valueStruct[float64]{V: Empty[float64]()},
			want: `{"v":null}`,
		},
		"empty value with present optional string": {
			arg:  valueStruct[string]{V: Of("")},
			want: `{"v":""}`,
		},
		"empty value with present optional int": {
			arg:  valueStruct[int]{V: Of(0)},
			want: `{"v":0}`,
		},
		"empty value with present optional float64": {
			arg:  valueStruct[float64]{V: Of(0.0)},
			want: `{"v":0}`,
		},
		"value with present optional string": {
			arg:  valueStruct[string]{V: Of("some string")},
			want: `{"v":"some string"}`,
		},
		"value with present optional int": {
			arg:  valueStruct[int]{V: Of(123)},
			want: `{"v":123}`,
		},
		"value with present optional float64": {
			arg:  valueStruct[float64]{V: Of(1.23)},
			want: `{"v":1.23}`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := string(bytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_MarshalJSON_StructWithPtr(t *testing.T) {
	tests := map[string]struct {
		arg     interface{}
		want    string
		wantErr bool
	}{
		"null with nil optional string": {
			arg:  ptrStruct[string]{V: nil},
			want: `{"v":null}`,
		},
		"null with nil optional int": {
			arg:  ptrStruct[int]{V: nil},
			want: `{"v":null}`,
		},
		"null with nil optional float64": {
			arg:  ptrStruct[float64]{V: nil},
			want: `{"v":null}`,
		},
		"null with empty optional string": {
			arg:  ptrStruct[string]{V: Empty[string]().Ptr()},
			want: `{"v":null}`,
		},
		"null with empty optional int": {
			arg:  ptrStruct[int]{V: Empty[int]().Ptr()},
			want: `{"v":null}`,
		},
		"null with empty optional float64": {
			arg:  ptrStruct[float64]{V: Empty[float64]().Ptr()},
			want: `{"v":null}`,
		},
		"empty value with present optional string": {
			arg:  ptrStruct[string]{V: Of("").Ptr()},
			want: `{"v":""}`,
		},
		"empty value with present optional int": {
			arg:  ptrStruct[int]{V: Of(0).Ptr()},
			want: `{"v":0}`,
		},
		"empty value with present optional float64": {
			arg:  ptrStruct[float64]{V: Of(0.0).Ptr()},
			want: `{"v":0}`,
		},
		"value with present optional string": {
			arg:  ptrStruct[string]{V: Of("some string").Ptr()},
			want: `{"v":"some string"}`,
		},
		"value with present optional int": {
			arg:  ptrStruct[int]{V: Of(123).Ptr()},
			want: `{"v":123}`,
		},
		"value with present optional float64": {
			arg:  ptrStruct[float64]{V: Of(1.23).Ptr()},
			want: `{"v":1.23}`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.arg)
			if err != nil {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := string(bytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_MarshalJSON_StructWithPtrAndOmitempty(t *testing.T) {
	tests := map[string]struct {
		arg     interface{}
		want    string
		wantErr bool
	}{
		"omitted with nil optional string": {
			arg:  ptrOmitStruct[string]{V: nil},
			want: `{}`,
		},
		"omitted with nil optional int": {
			arg:  ptrOmitStruct[int]{V: nil},
			want: `{}`,
		},
		"omitted with nil optional float64": {
			arg:  ptrOmitStruct[float64]{V: nil},
			want: `{}`,
		},
		"null with empty optional string": {
			arg:  ptrOmitStruct[string]{V: ptr(Empty[string]())},
			want: `{"v":null}`,
		},
		"null with empty optional int": {
			arg:  ptrOmitStruct[int]{V: ptr(Empty[int]())},
			want: `{"v":null}`,
		},
		"null with empty optional float64": {
			arg:  ptrOmitStruct[float64]{V: ptr(Empty[float64]())},
			want: `{"v":null}`,
		},
		"empty value with present optional string": {
			arg:  ptrOmitStruct[string]{V: ptr(Of(""))},
			want: `{"v":""}`,
		},
		"empty value with present optional int": {
			arg:  ptrOmitStruct[int]{V: ptr(Of(0))},
			want: `{"v":0}`,
		},
		"empty value with present optional float64": {
			arg:  ptrOmitStruct[float64]{V: ptr(Of(0.0))},
			want: `{"v":0}`,
		},
		"value with present optional string": {
			arg:  ptrOmitStruct[string]{V: ptr(Of("some string"))},
			want: `{"v":"some string"}`,
		},
		"value with present optional int": {
			arg:  ptrOmitStruct[int]{V: ptr(Of(123))},
			want: `{"v":123}`,
		},
		"value with present optional float64": {
			arg:  ptrOmitStruct[float64]{V: ptr(Of(1.23))},
			want: `{"v":1.23}`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.arg)
			if err != nil {
				t.Errorf("json.Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := string(bytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithValue_String(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    valueStruct[string]
		wantErr bool
	}{
		"omitted to empty optional": {
			arg:  `{}`,
			want: valueStruct[string]{V: Empty[string]()},
		},
		"null to empty optional": {
			arg:  `{"v":null}`,
			want: valueStruct[string]{V: Empty[string]()},
		},
		"empty value to present optional": {
			arg:  `{"v":""}`,
			want: valueStruct[string]{V: Of("")},
		},
		"value to present optional": {
			arg:  `{"v":"some string"}`,
			want: valueStruct[string]{V: Of("some string")},
		},
		"invalid value": {
			arg:     `{"v":123}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := valueStruct[string]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithValue_Int(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    valueStruct[int]
		wantErr bool
	}{
		"omitted to empty optional": {
			arg:  `{}`,
			want: valueStruct[int]{V: Empty[int]()},
		},
		"null to empty optional": {
			arg:  `{"v":null}`,
			want: valueStruct[int]{V: Empty[int]()},
		},
		"empty value to present optional": {
			arg:  `{"v":0}`,
			want: valueStruct[int]{V: Of(0)},
		},
		"value to present optional": {
			arg:  `{"v":123}`,
			want: valueStruct[int]{V: Of(123)},
		},
		"invalid value": {
			arg:     `{"v":"123"}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := valueStruct[int]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithValue_Float64(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    valueStruct[float64]
		wantErr bool
	}{
		"omitted to empty optional": {
			arg:  `{}`,
			want: valueStruct[float64]{V: Empty[float64]()},
		},
		"null to empty optional": {
			arg:  `{"v":null}`,
			want: valueStruct[float64]{V: Empty[float64]()},
		},
		"empty value to present optional": {
			arg:  `{"v":0}`,
			want: valueStruct[float64]{V: Of(0.0)},
		},
		"value to present optional": {
			arg:  `{"v":1.23}`,
			want: valueStruct[float64]{V: Of(1.23)},
		},
		"invalid value": {
			arg:     `{"v":"1.23"}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := valueStruct[float64]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithPtr_String(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    ptrStruct[string]
		wantErr bool
	}{
		"omitted to nil ptr": {
			arg:  `{}`,
			want: ptrStruct[string]{V: nil},
		},
		"null to nil ptr": {
			arg:  `{"v":null}`,
			want: ptrStruct[string]{V: nil},
		},
		"empty value to present optional ptr": {
			arg:  `{"v":""}`,
			want: ptrStruct[string]{V: ptr(Of(""))},
		},
		"value to present optional ptr": {
			arg:  `{"v":"some string"}`,
			want: ptrStruct[string]{V: ptr(Of("some string"))},
		},
		"invalid value": {
			arg:     `{"v":123}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ptrStruct[string]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithPtr_Int(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    ptrStruct[int]
		wantErr bool
	}{
		"omitted to nil ptr": {
			arg:  `{}`,
			want: ptrStruct[int]{V: nil},
		},
		"null to nil ptr": {
			arg:  `{"v":null}`,
			want: ptrStruct[int]{V: nil},
		},
		"empty value to present optional ptr": {
			arg:  `{"v":0}`,
			want: ptrStruct[int]{V: ptr(Of(0))},
		},
		"value to present optional ptr": {
			arg:  `{"v":123}`,
			want: ptrStruct[int]{V: ptr(Of(123))},
		},
		"invalid value": {
			arg:     `{"v":"some string"}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ptrStruct[int]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_UnmarshalJSON_StructWithPtr_Float64(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    ptrStruct[float64]
		wantErr bool
	}{
		"omitted to nil ptr": {
			arg:  `{}`,
			want: ptrStruct[float64]{V: nil},
		},
		"null to nil ptr": {
			arg:  `{"v":null}`,
			want: ptrStruct[float64]{V: nil},
		},
		"empty value to present optional ptr": {
			arg:  `{"v":0}`,
			want: ptrStruct[float64]{V: ptr(Of(0.0))},
		},
		"value to present optional ptr": {
			arg:  `{"v":1.23}`,
			want: ptrStruct[float64]{V: ptr(Of(1.23))},
		},
		"invalid value": {
			arg:     `{"v":"some string"}`,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ptrStruct[float64]{}
			err := json.Unmarshal([]byte(tt.arg), &got)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type valueStruct[T any] struct {
	V Optional[T] `json:"v"`
}

type ptrStruct[T any] struct {
	V *Optional[T] `json:"v"`
}

type ptrOmitStruct[T any] struct {
	V *Optional[T] `json:"v,omitempty"`
}

func ptr[T any](o Optional[T]) *Optional[T] {
	return &o
}
