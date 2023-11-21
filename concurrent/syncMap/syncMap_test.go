package syncmap

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestSyncMap(t *testing.T) {
	m := NewSyncMap()

	// 测试 Put 方法后 Get 方法是否能返回正确的值
	m.Put(1, 100)
	value, err := m.Get(1, 5*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 100, value)
	t.Log("测试 Put 后 Get 成功：", value)

	// 测试 Get 方法在键不存在时的阻塞行为
	go func() {
		time.Sleep(2 * time.Second)
		m.Put(2, 200)
		t.Log("值 200 已经被放入键 2")
	}()

	start := time.Now()
	value, err = m.Get(2, 5*time.Second)
	duration := time.Since(start)

	assert.NoError(t, err)
	assert.Equal(t, 200, value)
	assert.GreaterOrEqual(t, duration.Seconds(), 2.0, "Get 方法应该至少阻塞 2 秒钟")
	t.Logf("测试阻塞 Get：耗时 %v, 返回值 %v", duration, value)

	// 测试 Get 方法在超时后返回错误
	_, err = m.Get(3, 1*time.Second)
	assert.Error(t, err)
	t.Log("测试超时 Get 成功")
}
