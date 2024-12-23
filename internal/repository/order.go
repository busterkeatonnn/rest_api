package repository

import (
	"tspo_final/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(Order *models.Order) RepositoryResult {
	err := r.db.Save(Order).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: Order}
}

func (r *OrderRepository) FindAll() RepositoryResult {
	var Orders models.Orders

	err := r.db.Find(&Orders).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Orders}
}

func (r *OrderRepository) FindOneById(id uint) RepositoryResult {
	var Order models.Order

	err := r.db.Where(&models.Order{OrderID: id}).Take(&Order).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &Order}
}

func (r *OrderRepository) DeleteOneById(id uint) RepositoryResult {
	err := r.db.Delete(&models.Order{OrderID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *OrderRepository) DeleteByIds(ids *[]string) RepositoryResult {
	err := r.db.Where("OrderID IN (?)", *ids).Delete(&models.Orders{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
