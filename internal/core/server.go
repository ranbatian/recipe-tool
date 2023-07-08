package core

import (
	"github.com/gin-gonic/gin"
	"simon.com/recipe/internal/router"
)

func RunServer() {
	server := gin.Default()
	router.InitRouter(server)
	server.Run()
}
