package repositories

import (
	"go-hexagonal/core/domains"
	"go-hexagonal/core/ports"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	db.AutoMigrate(&domains.User{})
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(user *domains.User) error {
	return r.db.Create(&user).Error
}

func (r userRepositoryDB) GetAll() ([]domains.User, error) {
	users := []domains.User{}
	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r userRepositoryDB) GetByUserName(username string) (*domains.User, error) {
	user := domains.User{}
	tx := r.db.Where(&domains.User{Username: username}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
