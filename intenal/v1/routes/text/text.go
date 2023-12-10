package text

import (
	"github.com/gin-gonic/gin"
)

// PATH: api/text
func TextRouter(router gin.RouterGroup) {

	router.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "received",
		})
	})

}
