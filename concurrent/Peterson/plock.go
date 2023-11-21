package plock

// PetersonLock 分层锁
type PetersonLock struct {
	level   []int
	waiting []int
}

func NewPetersonLock(numThreads int) *PetersonLock {
	h := &PetersonLock{
		level:   make([]int, numThreads),
		waiting: make([]int, numThreads-1),
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
		h.level[i] = l
		h.waiting[l] = i

		for h.waiting[l] == i && h.existsOtherWithLevelGreaterOrEqual(i, l) {
		}
	}
}

// Unlock 释放锁
func (h *PetersonLock) Unlock(i int) {
	h.level[i] = -1
}

// existsOtherWithLevelGreaterOrEqual 检查是否存在其他进程在更高或相同的层级等待
func (h *PetersonLock) existsOtherWithLevelGreaterOrEqual(i, l int) bool {
	for k := 0; k < len(h.level); k++ {
		if k != i && h.level[k] >= l {
			return true
		}
	}
	return false
}
