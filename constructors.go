package optional

import (
	"github.com/mikkael131/go-optional/constraints"
	goConstraints "golang.org/x/exp/constraints"
)

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

func OfPtr[T any](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return Optional[T]{value: *value, present: true}
}

func OfEmpty[T constraints.Empty](value T) Optional[T] {
	if value.Empty() {
		return Empty[T]()
	}
	return Of(value)
}

func OfEmptyPtr[T constraints.Empty](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return OfEmpty(*value)
}

func OfGoEmpty[T goConstraints.Ordered](value T) Optional[T] {
	var zero T
	if zero == value {
		return Empty[T]()
	}
	return Of(value)
}

func OfGoEmptyPtr[T goConstraints.Ordered](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return OfGoEmpty(*value)
}
