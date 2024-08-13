package usecase

import (
	"go-api-commerce/model"
	"go-api-commerce/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

type UserListResponseApi struct {
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
	Success bool         `json:"success"`
}

type UserResponseApi struct {
	Message string     `json:"message"`
	Data    model.User `json:"data"`
	Success bool       `json:"success"`
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) GetUsers() UserListResponseApi {
	users, err := uc.repo.GetUsers()
	if err != nil {
		return UserListResponseApi{
			Message: "Error getting users",
			Data:    []model.User{},
			Success: false,
		}
	}

	return UserListResponseApi{
		Message: "Success getting users",
		Data:    users,
		Success: true,
	}
}

func (uc *UserUseCase) CreateUser(user model.User) UserResponseApi {
	foundUser, err := uc.repo.GetUserByEmail(user.Email)
	if err != nil {
		return UserResponseApi{
			Message: "Error checking user existence",
			Data:    model.User{},
			Success: false,
		}
	}
	if foundUser.ID != 0 {
		return UserResponseApi{
			Message: "User already exists",
			Data:    model.User{},
			Success: false,
		}
	}

	id, err := uc.repo.CreateUser(user)
	if err != nil {
		return UserResponseApi{
			Message: "Error creating user",
			Data:    model.User{},
			Success: false,
		}
	}

	user.ID = id
	return UserResponseApi{
		Message: "Success creating user",
		Data:    user,
		Success: true,
	}
}

func (uc *UserUseCase) GetUserByID(id int) UserResponseApi {
	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return UserResponseApi{
			Message: "Error getting, user is not found",
			Data:    model.User{},
			Success: false,
		}
	}

	return UserResponseApi{
		Message: "Success getting user",
		Data:    user,
		Success: true,
	}
}

func (uc *UserUseCase) UpdateUser(user model.User) UserResponseApi {
	existingUser, err := uc.repo.GetUserByID(user.ID)
	if err != nil {
		return UserResponseApi{
			Message: "Error getting, user is not found",
			Data:    model.User{},
			Success: false,
		}
	}

	// Verifique se o email foi alterado
	if existingUser.Email != user.Email {
		// Aqui você pode adicionar uma verificação de ACL no futuro
		return UserResponseApi{
			Message: "Error updating user, email cannot be changed",
			Data:    model.User{},
			Success: false,
		}
	}

	_, err = uc.repo.UpdateUser(user)
	if err != nil {
		return UserResponseApi{
			Message: "Error updating user",
			Data:    model.User{},
			Success: false,
		}
	}

	return UserResponseApi{
		Message: "Success updating user",
		Data:    user,
		Success: true,
	}
}

func (uc *UserUseCase) DeleteUser(id int) UserResponseApi {
	_, err := uc.repo.GetUserByID(id)
	if err != nil {
		return UserResponseApi{
			Message: "Error getting, user is not found",
			Data:    model.User{},
			Success: false,
		}
	}

	_, err = uc.repo.DeleteUser(id)
	if err != nil {
		return UserResponseApi{
			Message: "Error deleting user",
			Data:    model.User{},
			Success: false,
		}
	}

	return UserResponseApi{
		Message: "Success deleting user",
		Data:    model.User{},
		Success: true,
	}
}
