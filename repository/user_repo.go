package repository

import (
	"user-service/backend/model"
	"user-service/backend/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepoMethod struct {
	db  *gorm.DB
	log *logrus.Logger
}

func UserRepo(db *gorm.DB, log *logrus.Logger) *UserRepoMethod {
	return &UserRepoMethod{db, log}
}

func (r *UserRepoMethod) GetDataUserRepo(username string) (model.ResponseLogin, error) {
	r.log.Println("Execute function GetDataUserRepo")

	var user model.ResponseLogin
	err := r.db.Table(utils.TABEL_USER).
		Select(`user.id, user.username, emp.name, emp.email, role.role_name`).
		Joins(`LEFT JOIN employee emp ON user.id_employee = emp.id`).
		Joins(`LEFT JOIN role ON role.id = emp.id_role`).
		Where("user.deleted_at IS NULL").
		Find(&user).Error
	if err != nil {
		r.log.Error("Failed to get detail data user :", err)
		return model.ResponseLogin{}, err
	}

	return user, nil
}
