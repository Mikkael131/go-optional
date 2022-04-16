package optional

import (
	"reflect"
	"testing"
)

func TestOf(t *testing.T) {
	tests := map[string]struct {
		value string
		want  Optional[string]
	}{
		"empty string": {
			value: "",
			want:  Optional[string]{value: "", present: true},
		},
		"non blank string": {
			value: "some string",
			want:  Optional[string]{value: "some string", present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Of(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfPtr(t *testing.T) {
	someStr := "some string"
	tests := map[string]struct {
		value *string
		want  Optional[string]
	}{
		"nil": {
			value: nil,
			want:  Optional[string]{},
		},
		"string": {
			value: &someStr,
			want:  Optional[string]{value: someStr, present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := OfPtr(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testStruct struct {
	v string
}

func (t testStruct) Empty() bool {
	return t.v == ""
}

func TestOfEmpty(t *testing.T) {
	tests := map[string]struct {
		value testStruct
		want  Optional[testStruct]
	}{
		"empty struct": {
			value: testStruct{},
			want:  Optional[testStruct]{},
		},
		"non empty struct": {
			value: testStruct{v: "some string"},
			want:  Optional[testStruct]{value: testStruct{v: "some string"}, present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := OfEmpty(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfEmptyPtr(t *testing.T) {
	tests := map[string]struct {
		value *testStruct
		want  Optional[testStruct]
	}{
		"nil": {
			value: nil,
			want:  Optional[testStruct]{},
		},
		"empty struct ptr": {
			value: &testStruct{},
			want:  Optional[testStruct]{},
		},
		"non empty struct ptr": {
			value: &testStruct{v: "some string"},
			want:  Optional[testStruct]{value: testStruct{v: "some string"}, present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := OfEmptyPtr(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfEmptyPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfGoEmpty(t *testing.T) {
	tests := map[string]struct {
		value string
		want  Optional[string]
	}{
		"empty string": {
			value: "",
			want:  Optional[string]{},
		},
		"blank string": {
			value: "   ",
			want:  Optional[string]{value: "   ", present: true},
		},
		"non blank string": {
			value: "some string",
			want:  Optional[string]{value: "some string", present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := OfGoEmpty(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfGoEmptyPtr(t *testing.T) {
	emptyStr := ""
	blankStr := "   "
	someStr := "some string"
	tests := map[string]struct {
		value *string
		want  Optional[string]
	}{
		"empty string": {
			value: &emptyStr,
			want:  Optional[string]{},
		},
		"blank string": {
			value: &blankStr,
			want:  Optional[string]{value: blankStr, present: true},
		},
		"non blank string": {
			value: &someStr,
			want:  Optional[string]{value: someStr, present: true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := OfGoEmptyPtr(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OfGoEmptyPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
