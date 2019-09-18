package memoizer_go

import (
	"strconv"
	"time"
)

type ExpensiveFunction struct{}

func (self ExpensiveFunction) Compute(v string) int {
	ticker := time.NewTicker(3 * time.Second)
	<-ticker.C
	ticker.Stop()
	r, _ := strconv.Atoi(v)
	return r
}
