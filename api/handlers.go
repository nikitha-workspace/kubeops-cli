package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitha/kubeops-cli/cmd"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func VersionHandler(c *gin.Context) {

	cmd.Version()

	c.JSON(http.StatusOK, gin.H{
		"version": "1.0",
	})
}

func ValidateHandler(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "yaml file required",
		})
		return
	}

	path := "/tmp/" + file.Filename

	c.SaveUploadedFile(file, path)

	cmd.Validate(path)

	c.JSON(http.StatusOK, gin.H{
		"message": "Validation Successful",
	})
}

func DeployHandler(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "yaml required",
		})
		return
	}

	path := "/tmp/" + file.Filename

	c.SaveUploadedFile(file, path)

	cmd.Deploy(path)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deployment Started",
	})
}