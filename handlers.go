package ginx

import "github.com/gin-gonic/gin"

func HealthzHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func CommonRoute(rg *gin.RouterGroup) {
	rg.GET("/healthz", HealthzHandler)
}
