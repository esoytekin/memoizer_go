package memoizer_go

import (
	"log"
	"sync"
	"sync/atomic"

	future "github.com/esoytekin/go-future"
)

type Memoizer struct {
	e Computable
}

var cache sync.Map
var mutex sync.Mutex
var count uint64

func (self Memoizer) Compute(v string) int {

	result, ok := cache.Load(v)
	if !ok {

		mutex.Lock()
		result, ok = cache.Load(v)

		if !ok {
			ft := future.Start(func() (int, error) {
				atomic.AddUint64(&count, 1)
				return self.e.Compute(v), nil
			})
			cache.Store(v, ft)
			result = ft
		}

		mutex.Unlock()

	}

	r, err := result.(*future.FutureTask[int]).Get()

	if err != nil {
		log.Println(err)
		return 0
	}

	return r

}

func New(c Computable) Computable {
	return Memoizer{c}
}

func (self Memoizer) GetCount() uint64 {
	return count
}
