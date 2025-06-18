# 🚀 Go HTTP API — Profiling, Benchmarking & Performance Tuning

This project is a **high-performance HTTP API in Go**, designed to demonstrate how to **profile, benchmark, and optimize** backend services using Go’s native tools and concurrency patterns.

You’ll find:
- `pprof` profiling (CPU, memory, goroutines)
- Load testing with `hey` / `wrk`
- Object reuse with `sync.Pool`
- Buffered channels for queuing
- Worker goroutines with graceful shutdown
- Context-aware cancellation
- Clean Makefile workflows

---

## 🏗 Project Structure
├── handlers/ # HTTP handlers & worker pool
├── models/ # In-memory post store
├── main.go # App entrypoint with graceful shutdown
├── Makefile # Run tests, load test, profiles
├── go.mod
├── README.md
└── profiles/ # pprof output files


---

## 📦 Features

- 🧵 **Worker Pool**: Multiple goroutines process queued posts
- 💾 **sync.Pool**: Reduces allocations & GC pressure
- 🔄 **Buffered Channels**: Queue for incoming post writes
- 🛑 **Graceful Shutdown**: Drain and flush queue on SIGINT
- 🔍 **pprof Integration**: Inspect performance before exit
- 📊 **Load Tested**: With `hey` and `wrk`

---

## 🚀 Running the Project

```bash
go run main.go


