package syncmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSyncMap(t *testing.T) {
	m := NewSyncMap()

	// 测试 Put 方法后 Get 方法是否能返回正确的值（整数类型）
	m.Put(1, 100)
	value, err := m.Get(1, 5*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, 100, value)
	t.Log("测试 Put 后 Get 成功（整数）：", value)

	// 测试 Put 方法后 Get 方法是否能返回正确的值（字符串类型）
	m.Put(2, "hello")
	value, err = m.Get(2, 5*time.Second)
	assert.NoError(t, err)
	assert.Equal(t, "hello", value)
	t.Log("测试 Put 后 Get 成功（字符串）：", value)

	// 测试 Get 方法在键不存在时的阻塞行为
	go func() {
		time.Sleep(2 * time.Second)
		m.Put(3, []int{1, 2, 3})
		t.Log("切片 [1, 2, 3] 已经被放入键 3")
	}()

	start := time.Now()
	value, err = m.Get(3, 5*time.Second)
	duration := time.Since(start)

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, value)
	assert.GreaterOrEqual(t, duration.Seconds(), 2.0, "Get 方法应该至少阻塞 2 秒钟")
	t.Logf("测试阻塞 Get：耗时 %v, 返回值 %v", duration, value)

	// 测试 Get 方法在超时后返回错误
	_, err = m.Get(4, 1*time.Second)
	assert.Error(t, err)
	t.Log("测试超时 Get 成功")
	// 测试并发 Put 操作是否会尝试关闭已经关闭的通道
	t.Run("Concurrent Put Same Key", func(t *testing.T) {
		key := 5
		value := "concurrent test"

		// 启动多个协程进行并发的 Put 操作
		const numGoroutines = 10
		for i := 0; i < numGoroutines; i++ {
			go func() {
				m.Put(key, value)
			}()
		}

		// 等待一段时间以确保所有 Put 操作都有机会执行
		time.Sleep(100 * time.Millisecond)

		// 然后尝试获取该键的值
		retrievedValue, err := m.Get(key, 1*time.Second)

		// 验证获取的值和期望的值是否相同，以及没有发生错误
		assert.NoError(t, err)
		assert.Equal(t, value, retrievedValue)
		t.Log("并发 Put 同一个键测试成功")
	})
	//测试并发Get
	t.Run("Concurrent Get", func(t *testing.T) {
		key := 7
		value := "concurrent test"
		m.Put(key, value)

		// 启动多个协程进行并发的 Get 操作
		const numGoroutines = 3
		for i := 0; i < numGoroutines; i++ {
			go func() {
				m.Get(key, 1*time.Second)
			}()
		}

		// 等待一段时间以确保所有 Get 操作都有机会执行
		time.Sleep(3 * time.Second)

		// 然后尝试获取该键的值
		retrievedValue, err := m.Get(key, 1*time.Second)

		// 验证获取的值和期望的值是否相同，以及没有发生错误
		assert.NoError(t, err)
		assert.Equal(t, value, retrievedValue)
		t.Log("并发 Get 测试成功")
	})
}
