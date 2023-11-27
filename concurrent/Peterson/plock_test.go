package plock

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a = 0

func TestPLock(t *testing.T) {
	numThreads := 6
	var wg sync.WaitGroup
	lock := NewPetersonLock(numThreads)
	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				lock.Lock(i)
				a++
				lock.Unlock(i)
			}
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 600000, a)
}
