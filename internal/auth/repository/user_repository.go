package repository

import (
	"log"

	"github.com/mddsilva/gobet/internal/auth/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *dao.User) (dao.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	var error = u.db.Save((user))
	if error != nil {
		log.Print("Error on save user")
	}

	return *user, nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
