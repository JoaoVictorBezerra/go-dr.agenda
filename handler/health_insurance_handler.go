package handler

import (
	"dr.agenda/constants"
	"dr.agenda/dto"
	"dr.agenda/model"
	"dr.agenda/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthInsuranceHandler struct {
	useCase usecase.HealthInsuranceUseCase
}

func NewHealthInsuranceHandler(useCase usecase.HealthInsuranceUseCase) HealthInsuranceHandler {
	return HealthInsuranceHandler{
		useCase: useCase,
	}
}

func (handler *HealthInsuranceHandler) GetInsurances(ctx *gin.Context) {
	insurances, err := handler.useCase.GetActiveInsurances()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    nil,
			Message: constants.GetInsurancesError,
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response[[]model.HealthInsurance]{
		Data:    &insurances,
		Message: constants.GetInsuranceSuccess,
		Success: true,
	})
}
