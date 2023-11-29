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

func newEntry(v interface{}) *entry {
	return &entry{
		value: v,
		ready: make(chan struct{}),
	}
}

type SyncMap struct {
	mu   sync.Mutex
	data map[int]*entry
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		data: make(map[int]*entry),
	}
}

// Get 方法尝试获取键 k 对应的值。如果该键不存在，
// 它将阻塞直到 maxWaitingTime 时间，等待值被放入。
// 如果在超时时间内找到了值，它返回该值；否则返回错误。
func (m *SyncMap) Get(k int, maxWaitingTime time.Duration) (interface{}, error) {
	m.mu.Lock()
	ent, exists := m.data[k]
	if !exists {
		ent = &entry{ready: make(chan struct{})}
		m.data[k] = ent
	}
	m.mu.Unlock()

	select {
	case <-ent.ready:
		return ent.value, nil
	case <-time.After(maxWaitingTime):
		return nil, errors.New("")
	}
}

// Put 方法将键 k 的值设置为 v。如果该键对应的 entry 已经存在，
// 它更新该 entry 的值并关闭 ready 通道，通知等待该键的所有 Get 操作。
func (m *SyncMap) Put(k int, v interface{}) {
    m.mu.Lock()
    defer m.mu.Unlock()

    ent, exists := m.data[k]
    if !exists {
        ent = newEntry(v)
        m.data[k] = ent
    } else {
        // 如果 entry 已经存在，我们先更新值
        ent.value = v
    }

    // 检查 ready 通道是否已经关闭
    select {
    case <-ent.ready:
        return
    default:
        close(ent.ready)
    }
}

