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
		Success: false,
	})
}

func (handler *HealthInsuranceHandler) GetInsuranceById(ctx *gin.Context) {
	id := ctx.Param("id")

	insurance, err := handler.useCase.GetInsuranceById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    nil,
			Message: constants.GetInsurancesError,
			Success: false,
		})
		return
	}

	if insurance == nil {
		ctx.JSON(http.StatusNotFound, dto.Response[model.HealthInsurance]{
			Data:    nil,
			Message: constants.NotFoundInsurance,
			Success: true,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response[model.HealthInsurance]{
		Data:    insurance,
		Message: constants.GetInsuranceSuccess,
		Success: true,
	})
}

func (handler *HealthInsuranceHandler) CreateHealthInsurance(ctx *gin.Context) {
	var input dto.CreateHealthInsuranceRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			Data:    nil,
			Message: "Dados inválidos",
			Success: false,
		})
		return
	}

	insurance, err := handler.useCase.CreateInsurance(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    &err,
			Message: constants.CreateInsuranceError,
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response[model.HealthInsurance]{
		Data:    insurance,
		Message: constants.CreateInsuranceSuccess,
		Success: true,
	})
}
