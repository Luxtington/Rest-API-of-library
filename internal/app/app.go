package app

import (
	"ToGoList/internal/config"
	"ToGoList/internal/server"
	"log"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config, because of ", err)
	}

	appl := server.NewServer(cfg)
	log.Println("Server started at ", cfg.ServerHost)
	if err := appl.Run(); err != nil {
		log.Fatal("App run failed, because of ", err)
	}
}
