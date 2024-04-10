# 没有声明伪目标时，make会检查与目标同名文件是否存在，不存在或版本低则执行其下的命令。
# 声明伪目标后，就可以跳过这个检查过程，直接执行其下的命令
.PHONY: all build run gotool clean help

BINARY="main"

all: gotool build

# 在没有使用CGO的情况下，CGO_ENABLED=0可以让其静态编译；否则，需要设置为1，CGO的包会变成动态链接库，即生成的二进制文件对动态链接库有依赖
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"