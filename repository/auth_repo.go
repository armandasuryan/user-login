package repository

import (
	"auth/backend/model"
	"auth/backend/utils"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepoMethod struct {
	db  *gorm.DB
	log *logrus.Logger
}

func AuthRepo(db *gorm.DB, log *logrus.Logger) *AuthRepoMethod {
	return &AuthRepoMethod{db, log}
}

func (r *AuthRepoMethod) GetDataUserRepo(username string) (model.ResponseLogin, error) {
	r.log.Println("Execute function GetDataUserRepo")

	var user model.ResponseLogin
	if err := r.db.Table(utils.TABEL_USER).
		Select(`user.id, user.username, emp.name, emp.email, role.role_name`).
		Joins(`employee emp ON user.id_employee = emp.id`).
		Joins(`role ON role.id = emp.id_role`).
		Where("user.deleted_at IS NULL").
		Find(&user).Error; err != nil {
		r.log.Error("Failed to get detail data user :", err)
		return model.ResponseLogin{}, err
	}

	return user, nil
}

func (r *AuthRepoMethod) VerifyDataUserRepo(username, passwd string) string {

	var login model.Login

	// check if username exist
	if err := r.db.Table(utils.TABEL_USER).
		Where("username = ?", username).
		First(&login).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Error("User not found")
			return "User not found"
		}
		r.log.Error("Failed to get username in database")
		return "Failed to get username in database"
	}

	// verify password
	if !utils.VerifyPassword(passwd, login.Password) {
		r.log.Error("Password dosn't match")
		return "Password dosn't match"
	}

	return ""
}
