package handler

import (
	"dr.agenda/constants"
	"dr.agenda/dto"
	customErrors "dr.agenda/errors"
	"dr.agenda/model"
	"dr.agenda/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsuranceHandler struct {
	useCase usecase.InsuranceUseCase
}

func NewInsuranceHandler(useCase usecase.InsuranceUseCase) InsuranceHandler {
	return InsuranceHandler{
		useCase: useCase,
	}
}

func (handler *InsuranceHandler) GetInsurances(ctx *gin.Context) {
	insurances, err := handler.useCase.GetActiveInsurances()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    nil,
			Message: constants.GetInsurancesError,
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response[[]model.Insurance]{
		Data:    &insurances,
		Message: constants.GetInsuranceSuccess,
		Success: false,
	})
}

func (handler *InsuranceHandler) GetInsuranceById(ctx *gin.Context) {
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
		ctx.JSON(http.StatusNotFound, dto.Response[model.Insurance]{
			Data:    nil,
			Message: constants.NotFoundInsurance,
			Success: true,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response[model.Insurance]{
		Data:    insurance,
		Message: constants.GetInsuranceSuccess,
		Success: true,
	})
}

func (handler *InsuranceHandler) CreateInsurance(ctx *gin.Context) {
	var input dto.CreateInsuranceRequest

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

	ctx.JSON(http.StatusCreated, dto.Response[model.Insurance]{
		Data:    insurance,
		Message: constants.CreateInsuranceSuccess,
		Success: true,
	})
}

func (handler *InsuranceHandler) UpdateInsurance(ctx *gin.Context) {
	var input dto.UpdateInsuranceRequest
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		})
		return
	}

	insurance, err := handler.useCase.UpdateInsurance(id, input)
	if err != nil {
		if errors.Is(err, customErrors.InsuranceNotFoundError) {
			ctx.JSON(http.StatusNotFound, dto.Response[error]{
				Data:    nil,
				Message: constants.UpdateInsuranceError,
				Success: false,
			})
			return
		}

		if errors.Is(err, customErrors.UpdateInsuranceError) {
			ctx.JSON(http.StatusBadRequest, dto.Response[error]{
				Data:    nil,
				Message: constants.UpdateInsuranceError,
				Success: false,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    &err,
			Message: constants.UpdateInsuranceError,
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response[model.Insurance]{
		Data:    insurance,
		Message: constants.UpdateInsuranceSuccess,
		Success: true,
	})
}

func (handler *InsuranceHandler) SuspendInsurance(ctx *gin.Context) {
	id := ctx.Param("id")

	insurance, err := handler.useCase.SuspendHealthInsurance(id)
	if err != nil {
		if errors.Is(err, customErrors.InsuranceNotFoundError) {
			ctx.JSON(http.StatusNotFound, dto.Response[error]{
				Data:    nil,
				Message: constants.NotFoundInsurance,
				Success: false,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.Response[error]{
			Data:    &err,
			Message: constants.CreateInsuranceError,
			Success: false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response[model.Insurance]{
		Data:    insurance,
		Message: constants.CreateInsuranceSuccess,
		Success: true,
	})
}
