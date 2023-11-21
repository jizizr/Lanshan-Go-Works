```
➜  syncMap git:(main) ✗ go test -v

=== RUN   TestSyncMap
    syncMap_test.go:18: 测试 Put 后 Get 成功（整数）： 100
    syncMap_test.go:25: 测试 Put 后 Get 成功（字符串）： hello
    syncMap_test.go:31: 切片 [1, 2, 3] 已经被放入键 3
    syncMap_test.go:41: 测试阻塞 Get：耗时 2.001370291s, 返回值 [1 2 3]
    syncMap_test.go:46: 测试超时 Get 成功
--- PASS: TestSyncMap (3.00s)
PASS
ok      syncmap 3.277s
```