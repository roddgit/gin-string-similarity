package routes

import (
	"gin-string-similarity/controllers"

	"github.com/gin-gonic/gin"
)

//
// @Summary Compare name using jaro-winkler method
// @Accept  json
// @Produce  json
// @Param   request  body      payloads.CompareRequest  true  "CompareRequest JSON"
// @Success 200 {object}  payloads.CompareResponse
// @Router /compare-name [post]
func CompareRoute(router *gin.Engine) {
	router.POST("/api/compare-name", controllers.CompareHandler())

}
