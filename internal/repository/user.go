package repository

import (
	"tspo_final/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(User *models.User) RepositoryResult {
	err := r.db.Save(User).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: User}
}

func (r *UserRepository) FindAll() RepositoryResult {
	var Users models.Users

	err := r.db.Find(&Users).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Users}
}

func (r *UserRepository) FindOneById(id uint) RepositoryResult {
	var User models.User

	err := r.db.Where(&models.User{UserID: id}).Take(&User).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &User}
}

func (r *UserRepository) FindOneByUserName(username string) RepositoryResult {
	var User models.User

	err := r.db.Where(&models.User{Name: username}).Take(&User).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &User}
}

func (r *UserRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.User{UserID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *UserRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("UserID IN (?)", *ids).Delete(&models.Users{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
