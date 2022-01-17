package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sbswf/framework"
	"sbswf/framework/middleware"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	// core.Use(middleware.Timeout(1 * time.Second))

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8080",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("\nServer Shutdown: %s\n", err)
	}
}
