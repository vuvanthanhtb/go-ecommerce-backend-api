package controller

import (
	"github.com/gin-gonic/gin"
	s "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/service"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService s.IUserService
}

func NewUserController(userService s.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
