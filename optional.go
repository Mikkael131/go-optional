package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	value   T
	present bool
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.IsPresent()
}

func (o Optional[T]) IsPresent() bool {
	return o.present
}

func (o Optional[T]) IfPresent(consumer func(value T)) {
	if o.IsPresent() {
		consumer(o.value)
	}
}

func (o Optional[T]) IfPresentOrElse(consumer func(value T), runnable func()) {
	if o.IsPresent() {
		consumer(o.value)
	} else {
		runnable()
	}
}

func (o Optional[T]) Else(v T) T {
	if o.IsPresent() {
		return o.value
	}
	return v
}

func (o Optional[T]) ElseGet(supplier func() T) T {
	if o.IsPresent() {
		return o.value
	}
	return supplier()
}

func (o Optional[T]) ElseErr(err error) (T, error) {
	if o.IsPresent() {
		return o.value, nil
	}
	return o.value, err
}

func (o Optional[T]) ElseZero() T {
	if o.IsPresent() {
		return o.value
	}
	var zero T
	return zero
}

func (o Optional[T]) Filter(predicate func(v T) bool) Optional[T] {
	if o.IsPresent() && predicate(o.value) {
		return o
	}
	return Empty[T]()
}

func (o Optional[T]) Map(mapper func(v T) (r T, ok bool)) Optional[T] {
	if !o.IsPresent() {
		return Empty[T]()
	}
	v, ok := mapper(o.value)
	if !ok {
		return Empty[T]()
	}
	return Of(v)
}

func (o Optional[T]) FlatMap(mapper func(v T) Optional[T]) Optional[T] {
	if o.IsPresent() {
		return mapper(o.value)
	}
	return Empty[T]()
}

func (o Optional[T]) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("Optional[%v]", o.value)
	}
	return "Optional[empty]"
}

func (o Optional[T]) Ptr() *Optional[T] {
	if o.IsPresent() {
		return &o
	}
	return nil
}

func (o *Optional[T]) Val() Optional[T] {
	if o == nil {
		return Empty[T]()
	}
	return *o
}

var null = []byte("null")

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	value, ok := o.Get()
	if !ok {
		return null, nil
	}
	return json.Marshal(value)
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	if bytes.Compare(b, null) == 0 {
		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*o = Of(value)
	return nil
}
