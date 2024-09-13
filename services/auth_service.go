package services

import (
	"auth/backend/model"
	"auth/backend/repository"

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

	var response model.OTP
	verifyData := s.repo.VerifyDataUserRepo(payload.Username, payload.Password)

	return response, verifyData
}

func (s *AuthServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, error) {
	return s.repo.GetDataUserRepo(payload.Username)
}
