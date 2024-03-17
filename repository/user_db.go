package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return userRepository{db: db}
}

func (r userRepository) CreateUser(User) error {
	tx := r.db.Create(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r userRepository) GetByUsername(username string) (*User, error) {
	user := User{}
	tx := r.db.Model(User{Username: username}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
