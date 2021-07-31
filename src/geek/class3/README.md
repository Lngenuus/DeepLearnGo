# 作业内容

# 问题

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## 代码思路

参考kratos的设计:

- 对服务进行接口抽象用来满足同时启动多个服务的需要

- 使用WaitGroup的计数和等待机制确保全部的服务执行启动

- 全部的goroutine都是调用errorgroup中的GO方法创建

- 通过errorgroup启动的goroutine和上下文来优雅的退出全部goroutine

- 监听系统信号调用上下文取消优雅的退出

# [作业代码](./class3.go)

> 示例代码简单的启动了两个http服务监听在8080和8081端口 在根目录下通过以下命令启动demo:

```sh
go run main.go -geek class3
```
