package service

import (
	"fmt"
	"strconv"
	"time"

	r "github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/repo"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/utils/random"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/vuvanthanhtb/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email, purpose string) int
}

type userService struct {
	userRepo r.IUserRepository
	authRepo r.IAuthRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s\n", hashEmail)

	// 5. check OTP is available

	// 6. user spam

	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP
	otp := random.GenerateSixDigitOTP()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is::%d\n", otp)

	// 3. save OTP in Redis with expiration
	err := us.authRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. send email OTP
	err = sendto.SendTemplateEmailOTP([]string{email}, "xxx@gmail.com", "otp-auth.html", map[string]interface{}{"otp": strconv.Itoa(otp)})

	if err != nil {
		return response.ErrSendEmailOTP
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo r.IUserRepository, authRepo r.IAuthRepository) IUserService {
	return &userService{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}
