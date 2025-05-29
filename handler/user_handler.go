package handler

import (
	"dr.agenda/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return UserHandler{
		useCase: useCase,
	}
}

func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
