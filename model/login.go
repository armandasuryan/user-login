package model

type Login struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyOTP struct {
	UserName string `json:"user_name" validate:"required"`
	OTPCode  int    `json:"otp" validate:"required"`
}

type OTP struct {
	OTP      int    `json:"otp"`
	TimeLeft string `json:"time_left"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	RoleName string `json:"role_name"`
	Name     string `json:"employee_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
