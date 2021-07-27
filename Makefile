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

go-run-homework-optional:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	go run ./cmd/hw$${LESSON_NUMBER}/solution-hw-$${LESSON_NUMBER}.go -o

go-run-homework-exp:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	go run ./cmd/hw$${LESSON_NUMBER}/_experimental/exp-hw-$${LESSON_NUMBER}.go

go-run-homework-tests:
	@read -p "LESSON NUMBER: " LESSON_NUMBER; \
	cd ./homework/hw$${LESSON_NUMBER};\
	go test -v ./...

go-run-coverage:
	go test -count=1 -timeout=1s -short -race -covermode=atomic ./...

go-fmt:
	go fmt ./...


.PHONY: go-bench
#go-run-homework-001-bench1:
#	cd ./homework/hw001/helloWorldPrinter; \
#	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100000x -benchmem
##	go test -run=BenchmarkAddEmoji1 -bench=. -benchtime=100s

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
	firefox -new-tab "https://play.golang.org/p/mrsb_AtHflD"

ff-ref:
	firefox -new-tab "https://golang.org/ref/spec"

ff-wiki:
	firefox -new-tab "https://github.com/golang/go/wiki"

ff-tour:
	firefox -new-tab "https://tour.golang.org/welcome/1"

ff-spec:
	firefox -new-tab "https://golang.org/ref/spec#Type_identity"

ff-ref-go-pkg:
	firefox -new-tab "https://golang.org/pkg/"

ff-ref-go-modules:
	firefox -new-tab "https://golang.org/ref/mod#modules-overview"

ff-git-go-course:
	firefox -new-tab "https://github.com/rodkevich/go-course"

create-homework-dirs:
		@read -p "LESSON_NUMBER: " LESSON_NUMBER; \
		mkdir -p hw$${LESSON_NUMBER}/{cmd/hw$${LESSON_NUMBER}/lectures,internal,_experimental,pkg,api/hw$${LESSON_NUMBER}/v1,web,configs,init,build,test/testdata,test/hw$${LESSON_NUMBER},deploy,tools,docs/hw$${LESSON_NUMBER}}; \
		touch hw$${LESSON_NUMBER}/cmd/hw$${LESSON_NUMBER}/solution-hw-$${LESSON_NUMBER}.go; \
		touch hw$${LESSON_NUMBER}/_experimental/exp-hw-$${LESSON_NUMBER}.go; \
		touch hw$${LESSON_NUMBER}/test/hw$${LESSON_NUMBER}/hw-$${LESSON_NUMBER}_test.go; \
		touch hw$${LESSON_NUMBER}/docs/hw$${LESSON_NUMBER}/hw-$${LESSON_NUMBER}_notes.yml; \

lint:
	revive -formatter stylish ./... ;\
	golangci-lint run  ./...
