package factory

import (
	"database/sql"
	"dr.agenda/handler"
	"dr.agenda/repository"
	"dr.agenda/usecase"
)

func AuthenticationFactory(connection *sql.DB) handler.AuthenticationHandler {
	userRepositoryRepository := repository.NewUserRepository(connection)
	authUseCase := usecase.NewUserUseCase(userRepositoryRepository)
	return handler.NewAuthenticationHandler(authUseCase)
}
