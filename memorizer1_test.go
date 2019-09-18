package memoizer_go

import (
	"sync"
	"testing"

	"gotest.tools/assert"
)

func startTest(t *testing.T, c Computable) {
	testData := []string{"1", "1", "3", "3", "3", "5"}
	testDataEx := []int{1, 1, 3, 3, 3, 5}
	wg := sync.WaitGroup{}

	for i, td := range testData {
		ex := testDataEx[i]

		wg.Add(1)
		go func(td string, ex int) {
			defer wg.Done()
			res := c.Compute(td)
			assert.Assert(t, res == ex, "Expected %d, got %d", ex, res)
		}(td, ex)
	}
	wg.Wait()

}

func TestExpensiveFunc(t *testing.T) {

	l := new(ExpensiveFunction)
	startTest(t, l)
}

func TestMemoizer1(t *testing.T) {
	var e = ExpensiveFunction{}
	var m Computable = New(e)

	startTest(t, m)
	assert.Assert(t, m.(Memoizer).GetCount() == 3, "expected %d, got %d", 3, m.(Memoizer).GetCount())

}
