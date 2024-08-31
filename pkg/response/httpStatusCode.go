package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamInvalid  = 20003
	ErrInvalidToken      = 30001
	ErrInvalidOTP        = 30002
	ErrSendEmailOTP      = 30003
	ErrCodeUserHasExists = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Invalid token",
	ErrInvalidOTP:        "Invalid OTP",
	ErrSendEmailOTP:      "Failed to send OTP",
	ErrCodeUserHasExists: "User already exists",
}
