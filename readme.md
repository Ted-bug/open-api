# 1 项目结构
```txt
|--cmd
|--config           # 读取配置文件
|--internal
|   |--constants    # 常量配置
|   |--controller   # 控制器：业务逻辑
|   |--middleware   # 中间件
|   |--model        # 模型：定义数据库表对应的结构体
|   |--router       # 路由：配置控制器对应的路由
|   |--service      # 服务：功能点，按ddd模型分包
|   |--pkg         # 工具：数据库连接、redis连接、自定义辅助函数等第三方工具封装
|--main.go          # web入口文件
|--logs             # 日志
|--Dockerfile       # 容器配置文件
|--go.mod
|--go.sum
|--main.go          # 入口文件
|--Makefile         # 批处理脚本，make help 查看所有命令
|--README.md        # 说明文档
```
# 2 运行
1. 编译成可执行文件
```golang
go build -o open-api
```
2. 数据库初始化
```
open-api migrate
```
3. 运行
```
open-api run
```
