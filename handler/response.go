package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string){
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{
		"msg": msg,
		"errorCode": code,
	})
}
func sendSucess(ctx *gin.Context, op string, data interface{}, err error){
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data": data,
	})
}