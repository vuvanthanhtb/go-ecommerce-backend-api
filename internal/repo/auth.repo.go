package repo

import (
	"fmt"
	"time"

	"github.com/vuvanthanhtb/go-ecommerce-backend-api/global"
)

type IAuthRepository interface {
	AddOTP(email string, otp int, expiration int64) error
}

type authRepository struct {
}

// AddOTP implements IAuthRepository.
func (a *authRepository) AddOTP(email string, otp int, expiration int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expiration)).Err()
}

func NewAuthRepository() IAuthRepository {
	return &authRepository{}
}
