## 1 技术栈
以`Gin + Cobra + Viper + Gorm`搭建的api项目，包含路由、中间件、控制器、服务、model、工具等。

## 2 特性
- Makefile工具
- 构建docker镜像
- 命令行启动
- 配置热加载（资源类配置不适用）
- 数据迁移
- zap日志，按大小进行日志切割

## 3 项目结构
```txt
|--cmd
|--config           # 读取配置文件
|--internal
|   |--constants    # 全局通用常量配置
|   |--controller   # 业务层：参数校验、编写业务逻辑
|   |--middleware   # 中间件
|   |--model        # model层：定义数据库表对应的结构体
|   |--router       # 路由层：配置控制器对应的路由
|   |--service      # 服务层：通用功能点，按功能模块分包
|   |--pkg          # 工具：数据库连接、redis连接、自定义辅助函数等第三方工具封装
|   |--main.go      # web入口文件
|--logs             # 日志
|--Dockerfile       # 容器配置文件
|--go.mod
|--go.sum
|--main.go          # 入口文件
|--Makefile         # 批处理脚本，make help 查看所有命令
|--README.md        # 说明文档
```

## 4 运行
1. 编译成可执行文件
```golang
go build -o openapi ./main.go
// or make build
```
2. 数据库初始化
```
openapi migrate
```
3. 运行
```
openapi run
```
