package service

import r "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *r.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: r.NewUserRepo(),
	}
}

func (u *UserService) GetUserInfo() string {
	return u.userRepo.GetInfoUser()
}
