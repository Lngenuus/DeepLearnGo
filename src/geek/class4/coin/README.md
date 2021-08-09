# Kratos Project coin

```
    coin
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── api // 微服务使调用者的相关proto
    │   └── coin
    │       └── v1
    │           ├── coin.pb.go
    │           ├── coin.pb.validate.go
    │           ├── coin.proto
    │           ├── coin.swagger.json
    │           ├── coin_grpc.pb.go
    │           ├── coin_http.pb.go
    │           ├── error_reason.pb.go
    │           ├── error_reason.pb.validate.go
    │           ├── error_reason.proto
    │           └── error_reason.swagger.json
    ├── cmd // 硬币服务的启动入口
    │   └── coin
    │       ├── main.go
    │       ├── wire.go
    │       └── wire_gen.go
    ├── configs // 配置
    │   └── config.yaml
    ├── generate.go
    ├── go.mod
    ├── go.sum
    ├── internal // 服务内部代码
    │   ├── biz // 业务逻辑组装 do 定义的实现
    │   │   ├── README.md
    │   │   ├── biz.go
    │   │   └── coin.go
    │   ├── conf
    │   │   ├── conf.pb.go
    │   │   └── conf.proto
    │   ├── data // 访问各种数据库的封装 po的实现 biz调用该层组织内容
    │   │   ├── README.md
    │   │   ├── coin.go
    │   │   ├── data.go
    │   │   └── orm
    │   │       └── coin.go
    │   ├── server // http grpc swagger-ui 实例创建配置
    │   │   ├── grpc.go
    │   │   ├── http.go
    │   │   ├── server.go
    │   │   └── swagger.go
    │   └── service // 实现了api 定义的服务 调用了biz层内容生成dto
    │       ├── README.md
    │       ├── coin.go
    │       └── service.go
    └── third_party // api 依赖的第三方proto
        ├── README.md
        ├── errors
        │   └── errors.proto
        ├── google
        │   └── api
        │       ├── annotations.proto
        │       ├── http.proto
        │       └── httpbody.proto
        ├── protoc-gen-openapiv2
        │   └── options
        │       ├── annotations.proto
        │       └── openapiv2.proto
        └── validate
            ├── README.md
            └── validate.proto
```

