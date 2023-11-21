package nullable

type Nullable[T any] struct {
	value T
	valid bool
}

func Some[T any](value T) Nullable[T] {
	return Nullable[T]{value: value, valid: true}
}

func None[T any]() Nullable[T] {
	return Nullable[T]{valid: false}
}

func (n *Nullable[T]) IsValid() bool {
	return n.valid
}

func (n *Nullable[T]) Value() (T, bool) {
	return n.value, n.valid
}
