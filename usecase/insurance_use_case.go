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

func (useCase *InsuranceUseCase) GetActiveInsurances() ([]model.Insurance, error) {
	insurances, _ := useCase.repository.GetInsurances()
	var activeInsurancesList []model.Insurance
	activeInsurancesList = helpers.Filter(insurances, func(element model.Insurance) bool {
		return element.Status == enum.ACTIVE
	})
	return activeInsurancesList, nil
}

func (useCase *InsuranceUseCase) GetInsuranceById(id string) (*model.Insurance, error) {
	insurance, repositoryErr := useCase.repository.GetInsuranceById(id)

	if repositoryErr != nil {
		return nil, repositoryErr
	}

	if insurance == nil {
		return nil, customErrors.InsuranceNotFoundError
	}

	return insurance, nil
}

func (useCase *InsuranceUseCase) CreateInsurance(insuranceDTO dto.CreateInsuranceRequest) (*model.Insurance, error) {
	insurance, err := useCase.repository.CreateInsurance(insuranceDTO)
	return insurance, err
}

func (useCase *InsuranceUseCase) UpdateInsurance(id string, insuranceDTO dto.CreateInsuranceRequest) (*model.Insurance, error) {
	return nil, nil
}

func (useCase *InsuranceUseCase) SuspendHealthInsurance(id string) (*model.Insurance, error) {
	_, getInsuranceByIdErr := useCase.GetInsuranceById(id)

	if getInsuranceByIdErr != nil {
		return nil, getInsuranceByIdErr
	}

	insurance, err := useCase.repository.SuspendInsurance(id)

	if err != nil {
		return nil, err
	}

	return insurance, nil
}
