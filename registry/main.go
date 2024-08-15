package main

import (
	"fmt"
	"github.com/huiming23344/mindfs/registry/config"
	routers "github.com/huiming23344/mindfs/registry/router"
	"github.com/huiming23344/mindfs/registry/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}
	config.SetGlobalConfig(cfg)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	server.InitRegistryServer()
	go server.Check()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("Listen: %s\n", err)
	}
}
