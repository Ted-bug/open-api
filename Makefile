SCRIPT=go
BINARY=open-api
GOOS=linux
GOARCH=amd64
DATE=$(shell date +%Y%m%d)
VERSION=$(shell git describe --tags --always --dirty)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

.PHONY: all
# 格式化 Go 代码, 并编译生成二进制文件
all: gotool build

.PHONY: gotool
# 运行 Go 工具 'fmt' and 'vet'
gotool:
	$(SCRIPT) fmt ./...
	$(SCRIPT) vet ./...

.PHONY: build
# 编译 Go 代码, 生成二进制文件
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	$(SCRIPT) build \
	-ldflags "-X 'github.com/Ted-bug/open-api/cmd.Version=$(VERSION)' \
	-X 'github.com/Ted-bug/open-api/cmd.Branch=$(BRANCH)' \
	-X 'github.com/Ted-bug/open-api/cmd.Date=$(DATE)'" \
	-o ${BINARY} ./main.go

.PHONY: run
# 直接运行 Go 代码
run:
	$(SCRIPT) run ./... run

.PHONY: clean
# 移除二进制文件和 vim swap files
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: help
# 显示帮助信息
help:
	@echo ''
	@echo 'Usage:'
	@echo 'make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
    helpMessage = match(lastLine, /^# (.*)/); \
        if (helpMessage) { \
            helpCommand = substr($$1, 0, index($$1, ":")); \
            helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
            printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
        } \
    } \
    { lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help