package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	s "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/service"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *s.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: s.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// response.SuccessResponse(c, 20001, []string{"Thanh", "Phúc"})
	fmt.Println("My handler ------------->")
	response.ErrorResponse(c, response.ErrCodeParamInvalid)
}
