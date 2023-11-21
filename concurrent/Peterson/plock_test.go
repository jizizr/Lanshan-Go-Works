package plock

import (
	"fmt"
	"sync"
	"testing"
)

func TestPLock(t *testing.T) {
	numThreads := 100
	var wg sync.WaitGroup
	lock := NewPetersonLock(numThreads)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// fmt.Printf("进程 %d 正在尝试进入临界区...\n", id)
			lock.Lock(id)
			fmt.Printf("进程 %d 已进入临界区.\n", id)

			// 模拟临界区操作
			lock.Unlock(id)
			fmt.Printf("进程 %d 已退出临界区.\n", id)
		}(i)
	}

	wg.Wait()
}
