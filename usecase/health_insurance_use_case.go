package usecase

import (
	"dr.agenda/dto"
	"dr.agenda/enum"
	"dr.agenda/helpers"
	"dr.agenda/model"
	"dr.agenda/repository"
)

type HealthInsuranceUseCase struct {
	repository repository.HealthInsuranceRepository
}

func NewHealthInsuranceUseCase(repo repository.HealthInsuranceRepository) HealthInsuranceUseCase {
	return HealthInsuranceUseCase{
		repository: repo,
	}
}

func (useCase *HealthInsuranceUseCase) GetActiveInsurances() ([]model.HealthInsurance, error) {
	insurances, _ := useCase.repository.GetInsurances()
	var activeInsurancesList []model.HealthInsurance
	activeInsurancesList = helpers.Filter(insurances, func(element model.HealthInsurance) bool {
		return element.Status == enum.ACTIVE
	})
	return activeInsurancesList, nil
}

func (useCase *HealthInsuranceUseCase) GetInsuranceById(id string) (*model.HealthInsurance, error) {
	insurance, repositoryErr := useCase.repository.GetInsuranceById(id)

	if repositoryErr != nil {
		return nil, repositoryErr
	}

	return insurance, nil
}

func (useCase *HealthInsuranceUseCase) CreateInsurance(insuranceDTO dto.CreateHealthInsuranceRequest) (*model.HealthInsurance, error) {
	insurance, err := useCase.repository.CreateInsurance(insuranceDTO)
	return insurance, err
}

func (useCase *HealthInsuranceUseCase) UpdateInsurance(id int) (*model.HealthInsurance, error) {
	return nil, nil
}

func (useCase *HealthInsuranceUseCase) SuspendInsurance(id int) (*model.HealthInsurance, error) {
	return nil, nil
}
