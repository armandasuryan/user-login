package services

import (
	"auth/backend/model"
	"auth/backend/repository"
	"context"
	"crypto/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type AuthServiceMethod struct {
	repo  *repository.AuthRepoMethod
	redis *redis.Client
	log   *logrus.Logger
}

func AuthService(repo *repository.AuthRepoMethod, redis *redis.Client, log *logrus.Logger) *AuthServiceMethod {
	return &AuthServiceMethod{repo, redis, log}
}

func (s *AuthServiceMethod) LoginSvc(payload model.Login) (model.OTP, string) {
	s.log.Println("Execute function LoginSvc")

	var response model.OTP
	verifyData := s.repo.VerifyDataUserRepo(payload.Username, payload.Password)

	response.OTP, _ = s.GenerateOTPCode(6)
	err := s.redis.Set(context.TODO(), payload.Username+"_otp", response.OTP, time.Minute*5).Err() // time 5 minutes
	if err != nil {
		s.log.Error("Failed to store OTP in Redis:", err)
	}

	return response, verifyData
}

func (s *AuthServiceMethod) VerifyOTPSvc(payload model.VerifyOTP) (model.ResponseLogin, error) {
	s.log.Println("Execute function VerifyOTPSvc")

	storedOTP, err := s.redis.Get(context.TODO(), payload.Username+"_otp").Result()
	if err == redis.Nil {
		s.log.Error("OTP not found or expired")
		return model.ResponseLogin{}, err
	} else if err != nil {
		s.log.Error("Error getting OTP from Redis:", err)
		return model.ResponseLogin{}, err
	}

	storedOTPInt, _ := strconv.Atoi(storedOTP)

	if storedOTPInt != payload.OTPCode {
		s.log.Error("Invalid OTP")
		return model.ResponseLogin{}, err
	}

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
