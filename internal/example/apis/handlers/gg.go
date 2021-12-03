package handlers

import "github.com/gin-gonic/gin"

func GGHandler(c *gin.Context) {
	c.JSON(200, "gg")
}
