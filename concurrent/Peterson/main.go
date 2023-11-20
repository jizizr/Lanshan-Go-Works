package main

import (
	"fmt"
	"sync"
	"time"
)

type PetersonLock struct {
	flags  []bool // 标记
	victim int    // 轮转值
	num    int    // 总数
	lock   sync.Mutex
}

func NewPetersonLock(num int) *PetersonLock {
	return &PetersonLock{
		flags:  make([]bool, num),
		victim: 0,
		num:    num,
	}
}

func (p *PetersonLock) Lock(i int) {
	for j := 0; j < p.num; j++ {
		if j != i {
			p.flags[i] = true // 准备进入临界区
			p.victim = i      // 设置轮转值
			p.lock.Lock() //循环里用defer有可能出问题？
			for p.flags[j] && p.victim == i {
				// 忙等待(可以优化maybe)
			}
			p.lock.Unlock()
		}
	}
}

func (p *PetersonLock) Unlock(i int) {
	p.flags[i] = false
}

func main() {
	n := 20 // 假设有5个goroutine
	petersonLock := NewPetersonLock(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Printf("Goroutine %d: 尝试获取锁\n", i)
			petersonLock.Lock(i)
			// 临界区
			fmt.Printf("Goroutine %d: 在临界区\n", i)
			petersonLock.Unlock(i)
			fmt.Printf("Goroutine %d: 已释放锁\n", i)
		}(i)
	}

	// 等待一段时间以观察输出
	time.Sleep(50 * time.Second)
}
