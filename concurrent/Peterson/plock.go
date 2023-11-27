package plock

import (
	"sync/atomic"
)

// PetersonLock 分层锁
type PetersonLock struct {
	level   []int32
	waiting []int32
}

func NewPetersonLock(numThreads int) *PetersonLock {
	h := &PetersonLock{
		level:   make([]int32, numThreads),
		waiting: make([]int32, numThreads-1),
	}
	for i := range h.level {
		h.level[i] = -1
	}
	for i := range h.waiting {
		h.waiting[i] = -1
	}
	return h
}

// Lock 尝试获取锁
func (h *PetersonLock) Lock(i int) {
	for l := 0; l < len(h.waiting); l++ {
		atomic.StoreInt32(&h.level[i], int32(l))
		atomic.StoreInt32(&h.waiting[l], int32(i))

		for atomic.LoadInt32(&h.waiting[l]) == int32(i) && h.existsOtherWithLevelGreaterOrEqual(i, l) {
		}
	}
}

// Unlock 释放锁
func (h *PetersonLock) Unlock(i int) {
	atomic.StoreInt32(&h.level[i], -1)
}

// existsOtherWithLevelGreaterOrEqual 检查是否存在其他进程在更高或相同的层级等待
func (h *PetersonLock) existsOtherWithLevelGreaterOrEqual(i, l int) bool {
	for k := 0; k < len(h.level); k++ {
		if k != i && atomic.LoadInt32(&h.level[k]) >= int32(l) {
			return true
		}
	}
	return false
}
