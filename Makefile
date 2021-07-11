SHELL := /bin/bash
#COMMIT_TEXT := $(shell read -p "Enter: " enter ; echo $${enter})

#.PHONY: check
#check:
#ifdef COMMIT_TEXT
#	@echo "Warning: No message defined\; continue? [Y/n]"
#	@read line; if [ $$line = "n" ]; then echo aborting; exit 1 ; fi
#endif

.PHONY: cd
cd-zsh:
	@zsh

cd-bash:
	@bash

.PHONY: go
go-run-homework:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	go run ./cmd/hw$${LESSON_NUMBER}/solution-hw-$${LESSON_NUMBER}.go

go-run-homework-exp:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	go run ./cmd/hw$${LESSON_NUMBER}/_experimental/exp-hw-$${LESSON_NUMBER}.go

go-run-homework-optional:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	go run ./cmd/hw$${LESSON_NUMBER}/solution-hw-$${LESSON_NUMBER}.go -o

.PHONY: go-bench
#go-run-homework-001-bench1:
#	cd ./homework/hw001/helloWorldPrinter; \
#	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100000x -benchmem
##	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100s

.PHONY: go-test
#go-run-homework-001-bench2:
#	cd ./homework/hw001/helloWorldPrinter; \
#	go test -run=BenchmarkAddEmoji2 -bench=. -benchtime=100000x -benchmem

.PHONY: git
git:
	git status

git-commit-all:
	@read -p "COMMIT TEXT: " COMMIT_TEXT; \
	git add .; \
    git commit -m "$${COMMIT_TEXT}"

.PHONY: ff
ff-playground:
	firefox -new-tab "https://play.golang.org/p/n74anDvDqkA"

ff-tour:
	firefox -new-tab "https://tour.golang.org/welcome/1"

ff-ref-go-pkg:
	firefox -new-tab "https://golang.org/pkg/"

ff-ref-go-modules:
	firefox -new-tab "https://golang.org/ref/mod#modules-overview"

ff-git-go-course:
	firefox -new-tab "https://github.com/rodkevich/go-course"
