package syncmap

import (
	"errors"
	"sync"
	"time"
)

// entry 代表插入的数据结构
type entry struct {
	value interface{}
	//信号通道
	ready chan struct{}
}

func newEntry() *entry {
	return &entry{
		ready: make(chan struct{}),
	}
}

type SyncMap struct {
	data sync.Map
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

// Get 方法尝试获取键 k 对应的值。如果该键不存在，
// 它将阻塞直到 maxWaitingTime 时间，等待值被放入。
// 如果在超时时间内找到了值，它返回该值；否则返回错误。
func (m *SyncMap) Get(k int, maxWaitingTime time.Duration) (interface{}, error) {
	actual, _ := m.data.LoadOrStore(k, newEntry())
	ent := actual.(*entry)

	select {
	case <-ent.ready:
		return ent.value, nil
	case <-time.After(maxWaitingTime): //设置超时时间
		return 0, errors.New("")
	}
}

// Put 方法将键 k 的值设置为 v。如果该键对应的 entry 已经存在，
// 它更新该 entry 的值并关闭 ready 通道，通知等待该键的所有 Get 操作。
func (m *SyncMap) Put(k, v interface{}) {
	actual, loaded := m.data.LoadOrStore(k, &entry{value: v, ready: make(chan struct{})})
	ent := actual.(*entry)
	if loaded {
		ent.value = v
	}
	//关闭通道，通知所有阻塞的 Get 操作
	close(ent.ready)
}
