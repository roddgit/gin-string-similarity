package main

import (
	"gin-string-similarity/configs"
	_ "gin-string-similarity/docs"
	"gin-string-similarity/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var env = configs.EnvConfig()

// @title     gin-string-similarity
// @version         1.0
// @description     Compare name service API in Go using Gin framework.
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /api
func main() {
	// router := gin.Default()
	// routes.CompareRoute(router)

	router := gin.New()

	routes.CompareRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in
	_ = router.Run(":" + env["PORT"] + "")

}
