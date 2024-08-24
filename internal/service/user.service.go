package service

import (
	r "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/repo"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *r.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: r.NewUserRepo(),
// 	}
// }

// func (u *UserService) GetUserInfo() string {
// 	return u.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email, purpose string) int
}

type userService struct {
	userRepo r.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo r.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
