go-run-homework-001:
	go run ./cmd/hw001/solution-hw-001.go

go-run-homework-001-exp:
	go run ./cmd/hw001/_experimental/exp-hw-001.go

go-run-homework-002:
	go run ./cmd/hw002/solution-hw-002.go

go-run-homework-002-exp:
	go run ./cmd/hw002/solution-hw-002.go -e


go-run-homework-001-bench1:
	cd ./homework/hw001/hwp; \
	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100000x -benchmem
#	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100s

go-run-homework-001-bench2:
	cd ./homework/hw001/hwp; \
	go test -run=BenchmarkAddEmoji2 -bench=. -benchtime=100000x -benchmem

git-commit-all:
	git add .;\
    git commit -m "$m"
