// Package ast contains TypeScript's AST.
package ast

//go:generate go run ./genast.go -output ./ast_generated.go

type pool[T any] struct {
	data []T
}

func (p *pool[T]) allocate() *T {
	if len(p.data) == cap(p.data) {
		p.data = make([]T, 0, nextPoolSize(len(p.data)))
	}
	index := len(p.data)
	p.data = p.data[:index+1]
	return &p.data[index]
}

func nextPoolSize(size int) int {
	switch {
	case size < 16:
		return 16
	case size < 256:
		return size * 2
	}
	return size
}
