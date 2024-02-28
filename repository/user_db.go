package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{db: db}
}

func (r userRepository) CreateUser(User) (int64, error) {
	result := r.db.Create(&User{})
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected, nil
}

func (r userRepository) GetByUsername(username string) (User, error) {
	db_data := User{}
	result := r.db.Model(User{Username: username}).First(&db_data)
	if result.Error != nil {
		panic(result.Error)
	}
	return db_data, nil
}
