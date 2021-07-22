# 作业内容

#### 问题

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码

#### 答案

应该warp这个错误异常 因为我们需要知道详细的sql查询语句所执行的堆栈信息 直接打印错误可能会导致大量不需要的打印冗余错误占用资源 应抛出到业务逻辑层 由业务调用者决定是否打印信息 

数据操作层:

```golang
// 自定义错误内容 抛出异常
return "", xerrors.Wrapf(err, fmt.Sprintf("[queryName1ById] 信息查询失败, id=%v\n", id))

// 业务处理层选择是否打印错误 同时查看错误类型
// 错误详细信息输出
fmt.Printf("错误信息为:\n%+v\n", err)

// 根据业务需要执行对应的错误处理
if errors.Is(err, sql.ErrNoRows) {
    fmt.Println("\n\n业务侧对应的处理->没有查询到数据\n\n")
}
```

[代码例子如下](./class2.go)

> 该单例可以运行
