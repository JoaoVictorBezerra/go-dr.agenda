package factory

import (
	"database/sql"
	"dr.agenda/handler"
	"dr.agenda/repository"
	"dr.agenda/usecase"
)

func UserFactory(connection *sql.DB) handler.UserHandler {
	userRepository := repository.NewUserRepository(connection)
	userUseCase := usecase.NewUserUseCase(userRepository)
	return handler.NewUserHandler(userUseCase)
}
