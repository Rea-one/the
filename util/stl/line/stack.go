package line

type Stack[T any] struct {
	the    []T
	size   int
	cursor int
}
