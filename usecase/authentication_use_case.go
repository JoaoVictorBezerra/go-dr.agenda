package usecase

import "dr.agenda/repository"

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repository,
	}
}
