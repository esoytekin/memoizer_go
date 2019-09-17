package memoizer_go

import (
	"log"
	"sync"
	"sync/atomic"

	promise "github.com/fanliao/go-promise"
)

type Memoizer struct {
	E ExpensiveFunction
}

var cache sync.Map
var mutex sync.Mutex
var count uint64

func (self Memoizer) compute(v string) int {

	result, ok := cache.Load(v)
	if !ok {

		mutex.Lock()
		result, ok = cache.Load(v)

		if !ok {
			ft := promise.Start(func() (interface{}, error) {
				atomic.AddUint64(&count, 1)
				return self.E.compute(v), nil
			})
			cache.Store(v, ft)
			result = ft
		}

		mutex.Unlock()

	}

	r, err := result.(*promise.Future).Get()

	if err != nil {
		log.Println(err)
		return 0
	}

	return r.(int)

}
func (self Memoizer) GetCount() uint64 {
	return count
}
