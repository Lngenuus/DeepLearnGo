# 作业内容

# 问题

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## 代码思路

参考kratos的设计:

- 对服务的启动进行接口抽象用来满足同时启动多个服务的需要

- 通过waitGroup确保服务启动成功

- 通过errorGroup实现上下文的全部goroutine优雅退出

- 监听系统信号调用errorGroup生成的上下文

# [作业代码](./class3.go)

> 示例代码简单的启动了两个http服务监听在8080和8081端口 在根目录下通过以下命令启动demo:

```sh
go run main.go -geek class3
```
