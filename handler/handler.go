package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,gin.H{
		"msg": "POST opening",
	})
}
func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,gin.H{
		"msg": "GET opening",
	})
}
func DeleteOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,gin.H{
		"msg": "GET opening",
	})
}
func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,gin.H{
		"msg": "GET opening",
	})
}
func ListOpeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,gin.H{
		"msg": "GET opening",
	})
}