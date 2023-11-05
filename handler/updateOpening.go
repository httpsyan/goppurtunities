package handler

import (
	"net/http"

	"github.com/cplxx/goppurtunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context){
 request := updateOpeningRequest{}
 ctx.BindJSON(&request)
 if err := request.Validate(); err!= nil {
  logger.ErrorF("validation err %v", err.Error())
	sendError(ctx, http.StatusBadRequest,err.Error())
  return
 }

 id := ctx.Query("id")
 if id == ""{
    sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
    return
  }
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err!= nil {
    sendError(ctx, http.StatusNotFound,"opening not found")
    return 
  }

	if request.Role != "" {
		opening.Role = request.Role
	}
	
	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

 if request.Salary > 0 {
  opening.Salary = request.Salary
}

if err := db.Save(&opening).Error;err != nil {
	logger.ErrorF("error updating openin: %v", err)
	sendError(ctx, http.StatusInternalServerError, "error updating opening")
}
sendSucess(ctx, "update-opening", opening, nil)
}