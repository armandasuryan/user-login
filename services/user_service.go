package services

import (
	"user-service/backend/model"
	"user-service/backend/repository"

	"github.com/sirupsen/logrus"
)

type UserServiceMethod struct {
	repo *repository.UserRepoMethod
	log  *logrus.Logger
}

func UserService(repo *repository.UserRepoMethod, log *logrus.Logger) *UserServiceMethod {
	return &UserServiceMethod{repo, log}
}

func (s *UserServiceMethod) LoginUserSvc(payload model.Login) (model.OTP, error) {

	var response model.OTP

	return response, nil
}

func (s *UserServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, error) {
	return s.repo.GetDataUserRepo(payload.UserName)
}
