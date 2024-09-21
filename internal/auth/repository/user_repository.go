package repository

import (
	"errors"
	"log"

	"github.com/mddsilva/gobet/internal/auth/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *dao.User) (dao.User, error)
	FindByEmail(email string) (*dao.User, error)
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

func (u UserRepositoryImpl) FindByEmail(email string) (*dao.User, error) {
	user := &dao.User{}
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
