package main

import (
	"github.com/huiming23344/mindfs/client/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}
	config.SetGlobalConfig(cfg)

}
