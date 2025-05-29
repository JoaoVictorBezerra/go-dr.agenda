package factory

import (
	"database/sql"
	"dr.agenda/handler"
	"dr.agenda/repository"
	"dr.agenda/usecase"
)

func InsuranceFactory(connection *sql.DB) handler.InsuranceHandler {
	InsuranceRepository := repository.NewInsuranceRepository(connection)
	InsuranceUseCase := usecase.NewInsuranceUseCase(InsuranceRepository)
	return handler.NewInsuranceHandler(InsuranceUseCase)
}
