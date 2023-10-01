package main

import (
	"flag"
	"log"

	"github.com/ducnd58233/go-newsfeed-grpc/app/configs"
	"github.com/ducnd58233/go-newsfeed-grpc/app/internal/app"
	"github.com/ducnd58233/go-newsfeed-grpc/app/internal/app/service"
)

var path = flag.String("config", "config.yml", "config path for app service")

func main() {
	flag.Parse()
	conf, err := configs.GetAppConfig(*path)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	svc, err := service.NewService(conf)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	app.AppController{
		Service: *svc,
		Port:    conf.Port,
	}.Run()
}
