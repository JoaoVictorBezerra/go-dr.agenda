package usecase

import (
	"dr.agenda/dto"
	"dr.agenda/enum"
	customErrors "dr.agenda/errors"
	"dr.agenda/helpers"
	"dr.agenda/model"
	"dr.agenda/repository"
)

type InsuranceUseCase struct {
	repository repository.InsuranceRepository
}

func NewInsuranceUseCase(repo repository.InsuranceRepository) InsuranceUseCase {
	return InsuranceUseCase{
		repository: repo,
	}
}

func (uc *InsuranceUseCase) GetActiveInsurances() ([]model.Insurance, error) {
	insurances, _ := uc.repository.GetInsurances()
	var activeInsurancesList []model.Insurance
	activeInsurancesList = helpers.Filter(insurances, func(element model.Insurance) bool {
		return element.Status == enum.ACTIVE
	})
	return activeInsurancesList, nil
}

func (uc *InsuranceUseCase) GetInsuranceById(id string) (*model.Insurance, error) {
	insurance, repositoryErr := uc.repository.GetInsuranceById(id)

	if repositoryErr != nil {
		return nil, repositoryErr
	}

	if insurance == nil {
		return nil, customErrors.InsuranceNotFoundError
	}

	return insurance, nil
}

func (uc *InsuranceUseCase) CreateInsurance(insuranceDTO dto.CreateInsuranceRequest) (*model.Insurance, error) {
	insurance, err := uc.repository.CreateInsurance(insuranceDTO)
	return insurance, err
}

func (uc *InsuranceUseCase) UpdateInsurance(id string, dto dto.UpdateInsuranceRequest) (*model.Insurance, error) {
	_, getInsuranceByIdErr := uc.GetInsuranceById(id)

	if getInsuranceByIdErr != nil {
		return nil, getInsuranceByIdErr
	}

	updatedInsurance, updateInsuranceError := uc.repository.UpdateInsurance(id, dto)
	if updateInsuranceError != nil {
		return nil, customErrors.UpdateInsuranceError
	}

	return updatedInsurance, nil
}

func (uc *InsuranceUseCase) SuspendHealthInsurance(id string) (*model.Insurance, error) {
	_, getInsuranceByIdErr := uc.GetInsuranceById(id)

	if getInsuranceByIdErr != nil {
		return nil, getInsuranceByIdErr
	}

	insurance, err := uc.repository.SuspendInsurance(id)

	if err != nil {
		return nil, err
	}

	return insurance, nil
}
