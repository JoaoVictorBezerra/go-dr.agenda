package usecase

import (
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
