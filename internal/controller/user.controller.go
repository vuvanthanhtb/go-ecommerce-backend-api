package controller

import (
	"github.com/gin-gonic/gin"
	s "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/service"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/vo"
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
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid)
		return
	}
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
