package factory

import (
	"database/sql"
	"dr.agenda/handler"
	"dr.agenda/repository"
	"dr.agenda/usecase"
)

func HealthInsuranceFactory(connection *sql.DB) handler.HealthInsuranceHandler {
	HealthInsuranceRepository := repository.NewHealthInsuranceRepository(connection)
	HealthInsuranceUseCase := usecase.NewHealthInsuranceUseCase(HealthInsuranceRepository)
	return handler.NewHealthInsuranceHandler(HealthInsuranceUseCase)
}
