package optional

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		arg        func(v string) (int, bool)
		want       Optional[int]
		wantCalled bool
	}{
		"do not map empty optional": {
			optional:   Empty[string](),
			arg:        nil,
			want:       Empty[int](),
			wantCalled: false,
		},
		"map present optional with returned ok": {
			optional: Of("123"),
			arg: func(v string) (int, bool) {
				s, err := strconv.Atoi(v)
				return s, err == nil
			},
			want:       Of(123),
			wantCalled: true,
		},
		"do not map present optional with returned nok": {
			optional: Of("this is not a number"),
			arg: func(v string) (int, bool) {
				s, err := strconv.Atoi(v)
				return s, err == nil
			},
			want:       Empty[int](),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			fnWrapper := func(v string) (int, bool) {
				called = true
				return tt.arg(v)
			}

			got := Map(tt.optional, fnWrapper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
			if called != tt.wantCalled {
				t.Errorf("Map() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	tests := map[string]struct {
		optional   Optional[string]
		arg        func(v string) Optional[int]
		want       Optional[int]
		wantCalled bool
	}{
		"do not flat map empty optional": {
			optional:   Empty[string](),
			arg:        nil,
			want:       Empty[int](),
			wantCalled: false,
		},
		"flat map with returned present optional": {
			optional: Of("123"),
			arg: func(v string) Optional[int] {
				return Of(123)
			},
			want:       Of(123),
			wantCalled: true,
		},
		"flat map with returned empty optional": {
			optional: Of("123"),
			arg: func(v string) Optional[int] {
				return Empty[int]()
			},
			want:       Empty[int](),
			wantCalled: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			fnWrapper := func(v string) Optional[int] {
				called = true
				return tt.arg(v)
			}

			got := FlatMap(tt.optional, fnWrapper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
			if called != tt.wantCalled {
				t.Errorf("FlatMap() fn called = %v, want %v", called, tt.wantCalled)
			}
		})
	}
}
