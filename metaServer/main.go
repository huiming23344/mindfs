package main

import (
	"fmt"
	"github.com/huiming23344/mindfs/metaServer/apis"
	"github.com/huiming23344/mindfs/metaServer/config"
	routers "github.com/huiming23344/mindfs/metaServer/router"
	"github.com/huiming23344/mindfs/metaServer/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}
	config.SetGlobalConfig(cfg)

	// 创建一个信号通道
	sigChan := make(chan os.Signal, 1)

	// 注册信号通道，以接收 os.Interrupt 和 syscall.SIGTERM
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// 启动一个 goroutine 来监听信号
	go func() {
		<-sigChan // 等待接收信号
		apis.Unregister()
		os.Exit(0) // 退出程序
	}()

	router := routers.InitRouter()
	server.InitServer()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	apis.Register()
	go apis.Heartbeat()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Listen: %s\n", err)
	}
}
