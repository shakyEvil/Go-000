按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。


└── demo
    ├── api
    │   ├── api.go
    │   ├── api.proto
    │   └── client.go
    ├── cmd
    │   ├── cmd
    │   └── main.go
    ├── configs
    │   ├── db.yaml
    │   ├── grpc.yaml
    │   └── redis.yaml
    ├── go.mod
    └── internal
        ├── dao
        │   └── dao.go
        ├── model
        │   └── model.go
        ├── server
        │   └── server.go
        └── service
            └── service.go
