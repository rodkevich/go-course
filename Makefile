SHELL := /bin/bash
COMMIT_TEXT := $(shell read -p "Enter: " enter ; echo $${enter})

.PHONY: check
check:
ifndef COMMIT_TEXT
	@echo "Warning: No message defined\; continue? [Y/n]"
	@read line; if [ $$line = "n" ]; then echo aborting; exit 1 ; fi
endif

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

git-commit-all: check
	git add .;\
    git commit -m "${COMMIT_TEXT}"

playground:
	firefox -new-tab "https://play.golang.org/"
