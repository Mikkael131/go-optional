package optional

func Map[T, R any](o Optional[T], mapper func(o T) (result R, ok bool)) Optional[R] {
	value, ok := o.Get()
	if !ok {
		return Empty[R]()
	}
	v, ok := mapper(value)
	if !ok {
		return Empty[R]()
	}
	return Of(v)
}

func FlatMap[T, R any](o Optional[T], mapper func(o T) Optional[R]) Optional[R] {
	value, ok := o.Get()
	if !ok {
		return Empty[R]()
	}
	return mapper(value)
}
