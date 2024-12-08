.PHONY: test
.PHONY: all
.PHONY: clean
.PHONY: build
.PHONY: run
build:
        @go build main.go

run:    build
        @./main 
