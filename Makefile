run:
	go run main.go
build:
	go build -o api-server main.go

loadtest-hey:
	hey -n 10000 -c 100 http://localhost:8070/posts

loadtest-wrk:
	wrk -t12 -c100 -d30s http://localhost:8070/posts

profile-cpu:
	curl -o cpu.pprof http://localhost:6060/debug/pprof/profile?seconds=30
	go tool pprof -http=:7070 api-server cpu.pprof

profile-mem:
	curl -o mem.pprof http://localhost:6060/debug/pprof/heap
	go tool pprof -http=:7070 api-server mem.pprof

benchmark:
	go test -bench=. -benchmem ./handlers > benchtest.txt
	@echo "Benchmark resutls savet to bench_latest.txt"

benchmark-compare:
	benchstat bench_old.txt bench_latest.txt
