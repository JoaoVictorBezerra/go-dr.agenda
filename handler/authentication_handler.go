package handler

import (
	"dr.agenda/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationHandler struct {
	useCase usecase.UserUseCase
}

func NewAuthenticationHandler(uc usecase.UserUseCase) AuthenticationHandler {
	return AuthenticationHandler{
		useCase: uc,
	}
}

func (h *AuthenticationHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
