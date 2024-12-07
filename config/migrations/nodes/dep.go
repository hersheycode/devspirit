package schema

type Collection[T any] struct {
	Group []T
	Seen  map[string]T
}
