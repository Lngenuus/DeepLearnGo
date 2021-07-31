package class3

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// 参考kratos的抽象和实现
// Server 启动服务的抽象 可以是http/grpc 一下示例仅用http启动服务
type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// HTTP 接口的实现
type HttpServer struct {
	*http.Server
}

// Start 启动 http server
func (hs *HttpServer) Start(ctx context.Context) error {
	return hs.ListenAndServe()
}

// Stop 停止http服务
func (hs *HttpServer) Stop(ctx context.Context) error {
	return hs.Shutdown(ctx)
}

type appKey struct{}

type App struct {
	// svrs 可以启用多个http server
	svrs []Server
	ctx  context.Context
	// 传入上下文取消方法 用来优雅的退出
	cancel func()
}

// Run 启动应用
func (a *App) Run() error {
	// 局部变量作为key来确保全局唯一标识
	ctx := context.WithValue(a.ctx, appKey{}, a)

	// errorgroup 的特性确保全部的 goroutine 退出
	eg, ctx := errgroup.WithContext(ctx)

	// waitgroup 用来确保全部的服务正常启动
	wg := sync.WaitGroup{}
	for _, svr := range a.svrs {
		// 赋值问题
		s := svr
		eg.Go(func() error {
			<-ctx.Done()
			return s.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return s.Start(ctx)
		})
	}
	wg.Wait()
	fmt.Println("服务启动成功！！！！")

	// 停止信号
	sig := make(chan os.Signal)

	// Go 不允许监听 SIGKILL/SIGSTOP 信号
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case s := <-sig:
			fmt.Printf("\n\n[%v]退出中...\n\n", s)
			return a.Quit()
		}
	})
	// 当有错误返回且错误为上下文取消的时候意味着是接收到退出信号 否则抛出异常错误
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Quit 优雅退出应用
func (a *App) Quit() error {
	if a.cancel != nil {
		// 实际调用根节上下文取消函数
		a.cancel()
	}
	fmt.Println("已经安全退出!")
	return nil
}

// NewApp 简单的demo启动创建
func NewApp() {

	//  启动两个http server
	mux1 := http.NewServeMux()
	mux2 := http.NewServeMux()

	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HTTP SERVER 1:\n    Hello, %q\n", html.EscapeString(r.URL.Path))
	})
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HTTP SERVER 2:\n    Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	hs1 := &HttpServer{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: mux1,
		},
	}
	hs2 := &HttpServer{
		Server: &http.Server{
			Addr:    ":8081",
			Handler: mux2,
		},
	}
	// 生成含有撤销机制的上下文
	ctx, cancel := context.WithCancel(context.Background())
	app := &App{
		svrs:   []Server{hs1, hs2},
		ctx:    ctx,
		cancel: cancel,
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
