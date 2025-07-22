package repository

import (
	"subscription-service/model"

	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepo) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.db.Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepo) GetByID(id int) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, id).Error
	return &sub, err
}

func (r *SubscriptionRepo) Update(id int, data *model.Subscription) error {
	return r.db.Model(&model.Subscription{}).Where("id = ?", id).Updates(data).Error
}

func (r *SubscriptionRepo) Delete(id int) error {
	return r.db.Delete(&model.Subscription{}, id).Error
}

func (r *SubscriptionRepo) GetTotal(service string, userId string) (float64, error) {
	var total float64
	query := r.db.Model(&model.Subscription{}).Select("SUM(price)")
	if service != "" {
		query = query.Where("service = ?", service)
	}
	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}
	err := query.Scan(&total).Error
	return total, err
}
