```
➜  syncMap git:(main) ✗ go test -v

=== RUN   TestSyncMap
    syncMap_test.go:17: 测试 Put 后 Get 成功： 100
    syncMap_test.go:23: 值 200 已经被放入键 2
    syncMap_test.go:33: 测试阻塞 Get：耗时 2.001265417s, 返回值 200
    syncMap_test.go:38: 测试超时 Get 成功
--- PASS: TestSyncMap (3.00s)
PASS
ok      syncmap 3.127s
```