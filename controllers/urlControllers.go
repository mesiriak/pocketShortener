package controllers

import "github.com/gin-gonic/gin"

func UrlFetch(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"url": "Enter now :)"})
}
