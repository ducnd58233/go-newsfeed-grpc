package app

import (
	"fmt"

	"github.com/ducnd58233/go-newsfeed-grpc/app/internal/app/service"
	"github.com/gin-gonic/gin"
)

type AppController struct {
	service.Service
	Port int
}

func (c AppController) Run() {
	r := gin.Default()

	// v1Router := r.Group("/api/v1")
	r.Run(fmt.Sprintf(":%d", c.Port))
}