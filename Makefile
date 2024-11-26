BINARY="open-api"

.PHONY: all
# 发布操作：格式化、编译
all: gotool build

.PHONY: gotool
# 格式化代码
gotool:
	go fmt ./
	go vet ./

.PHONY: build
# 编译项目
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'github.com/Ted-bug/open-api/cmd.Version=$(git describe --tags --always --dirty)' -X 'github.com/Ted-bug/open-api/cmd.Branch=$(git rev-parse --abbrev-ref HEAD)' -X 'github.com/Ted-bug/open-api/cmd.Date=$(date)'" -o ${BINARY}

.PHONY: run
# 直接运行项目
run:
	go run ./

.PHONY: clean
# 清理编译等生成的文件
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: help
# 显示帮助信息
help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
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