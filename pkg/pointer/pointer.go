package pointer

func Val[T any](p *T) T {
	if p == nil {
		p = new(T)
	}
	return *p
}

func New[T any](o T) *T {
	return &o
}
