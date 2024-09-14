package services

import (
	"auth/backend/model"
	"auth/backend/repository"
	"crypto/rand"
	"strconv"

	"github.com/sirupsen/logrus"
)

type AuthServiceMethod struct {
	repo *repository.AuthRepoMethod
	log  *logrus.Logger
}

func AuthService(repo *repository.AuthRepoMethod, log *logrus.Logger) *AuthServiceMethod {
	return &AuthServiceMethod{repo, log}
}

func (s *AuthServiceMethod) LoginSvc(payload model.Login) (model.OTP, string) {
	s.log.Println("Execute function LoginSvc")

	var response model.OTP
	verifyData := s.repo.VerifyDataUserRepo(payload.Username, payload.Password)

	response.OTP, _ = s.GenerateOTPCode(6)
	return response, verifyData
}

func (s *AuthServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, error) {
	s.log.Println("Execute function VerifyOTPSvc")
	return s.repo.GetDataUserRepo(payload.Username)
}

func (s *AuthServiceMethod) GenerateOTPCode(maxDigit int) (int, error) {
	s.log.Println("Execute function GenerateOTPCode")

	codes := make([]byte, maxDigit) // set digit
	if _, err := rand.Read(codes); err != nil {
		s.log.Error("Failed get random code")
		return 0, err
	}

	for i := 0; i < maxDigit; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	result, _ := strconv.Atoi(string(codes))
	return result, nil
}
