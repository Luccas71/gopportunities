package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHanlder(ctx *gin.Context) {
	// request := struct {
	// 	Role string `json:"role"`
	// }{} // definindo um struct anônimo e iniciando ele vazio

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request) //pega o que vem da request e joga pra um formato definido

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
	}
}
