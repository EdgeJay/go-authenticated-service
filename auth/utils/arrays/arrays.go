package arrays

type Iterator[M any] struct {
	index       int
	middlewares []M
}

func NewIterator[M any](middlewares []M) (it *Iterator[M]) {
	it = &Iterator[M]{
		index:       0,
		middlewares: middlewares,
	}
	return
}

func (it *Iterator[M]) NextFunc() (m M, done bool) {
	if it.index >= len(it.middlewares) {
		done = true
		return
	}
	m = it.middlewares[it.index]
	it.index += 1
	return
}
