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

func (s *AuthServiceMethod) LoginSvc(payload model.Login) (model.OTP, error) {

	var response model.OTP

	return response, nil
}

func (s *AuthServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, error) {
	return s.repo.GetDataUserRepo(payload.UserName)
}
