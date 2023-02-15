package pocketShortener

import (
	"github.com/gin-gonic/gin"
	"pocketShortener/controllers"
)

func main() {
	app := gin.Default()

	app.GET("/", controllers.UrlFetch)
}
