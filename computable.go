package memoizer_go

type Computable interface {
	Compute(v string) int
}
