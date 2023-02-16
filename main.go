package main

import (
	"github.com/gin-gonic/gin"
	"pocketShortener/controllers"
	"pocketShortener/inits"
)

func init() {
	inits.LoadEnvs()
	inits.ConnectToDB()
}

func main() {
	app := gin.Default()

	app.GET("/", controllers.UrlFetch)
}
