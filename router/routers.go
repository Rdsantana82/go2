package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouters(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	v1.GET("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET Opening",
		})
	})

	v1.POST("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "POST Opening",
		})
	})

	v1.PUT("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "PUT Opening",
		})
	})

	v1.DELETE("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "DELETE Opening",
		})
	})

	v1.GET("/openings", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET Openings",
		})
	})
}
