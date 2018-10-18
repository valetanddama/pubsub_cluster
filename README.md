### Запустить сервер
```go
go run pkg/cmd/cluster/main.go
go run pkg/cmd/cluster/main.go --getErrors
```

### Собрать билд
```bash
go build -o ./bin/cluster pkg/cmd/cluster/*.go
```

### Запустить бенчмарки
```bash
go test -benchmem -bench=. test/benchmark_test.go

BenchmarkSendMessageForRoleGenerator-8             30000             51225 ns/op             240 B/op          7 allocs/op
BenchmarkGetMessageForRoleHandler-8                20000             55303 ns/op             516 B/op         15 allocs/op
BenchmarkCheckServerRole-8                         10000            150301 ns/op             656 B/op         20 allocs/op
BenchmarkGetRedisConn-8                         1000000000               2.56 ns/op            0 B/op          0 allocs/op
BenchmarkPing-8                                    30000             48672 ns/op             148 B/op          5 allocs/op
```
