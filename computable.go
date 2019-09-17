package memoizer_go

type Computable interface {
	compute(v string) int
}
