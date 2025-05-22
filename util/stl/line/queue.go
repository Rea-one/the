package line

type Queue[T any] struct {
	the    []T
	size   int
	cursor int
}

func NewQueue[T any]() *Queue[T] {
	tar := &Queue[T]{}
	tar.the = make([]T, 0)
	tar.size = 0
	tar.cursor = -1
	return tar
}

func (tar *Queue[T]) Push(item T) {
	tar.size++
	tar.cursor += 1
	tar.the = append(tar.the, item)
}

func (tar *Queue[T]) Pop() T {
	if tar.size == 0 {
		var empty T
		return empty
	}
	tar.size--
	tar.cursor -= 1
	return tar.the[]
}
