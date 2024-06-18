package handler

import (
	"net/http"

	"github.com/Luccas71/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
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

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
	}

	sendSuccess(ctx, "create-opening", opening)
}
