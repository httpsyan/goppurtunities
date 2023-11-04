package handler

import (
	"net/http"

	"github.com/cplxx/goppurtunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
    Company: request.Company,
    Location: request.Location,
		Remote: *request.Remote,
    Link: request.Link,
    Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err!=nil {
		logger.ErrorF("error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creatomg opening on databse")
		return
	}

	sendSucess(ctx, "create-opening", opening, nil)
}