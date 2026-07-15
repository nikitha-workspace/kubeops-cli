package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitha/kubeops-cli/api"
)

func main() {

	r := gin.Default()

	r.GET("/health", api.HealthHandler)
	r.GET("/version", api.VersionHandler)
	r.POST("/validate", api.ValidateHandler)
	r.POST("/deploy", api.DeployHandler)

	r.Run(":8081")
}