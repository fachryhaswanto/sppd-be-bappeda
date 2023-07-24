package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type BidangRepository interface {
	FindAll() ([]model.Bidang, error)
	FindById(id int) (model.Bidang, error)
	Create(bidang model.Bidang) (model.Bidang, error)
	Update(bidang model.Bidang) (model.Bidang, error)
	Delete(bidang model.Bidang) (model.Bidang, error)
}

type bidangRepository struct {
	db *gorm.DB
}

func NewBidangRepository(db *gorm.DB) *bidangRepository {
	return &bidangRepository{db}
}

func (r *bidangRepository) FindAll() ([]model.Bidang, error) {
	var bidangs []model.Bidang

	var err = r.db.Where("id not in (8)").Find(&bidangs).Error

	return bidangs, err
}

func (r *bidangRepository) FindById(id int) (model.Bidang, error) {
	var bidang model.Bidang

	var err = r.db.Where("id not in (8)").Find(&bidang, id).Error

	return bidang, err
}

func (r *bidangRepository) Create(bidang model.Bidang) (model.Bidang, error) {
	var err = r.db.Create(&bidang).Error

	return bidang, err
}

func (r *bidangRepository) Update(bidang model.Bidang) (model.Bidang, error) {
	var err = r.db.Model(&bidang).Updates(model.Bidang{Nama_Bidang: bidang.Nama_Bidang, Singkatan: bidang.Singkatan}).Error

	return bidang, err
}

func (r *bidangRepository) Delete(bidang model.Bidang) (model.Bidang, error) {
	var err = r.db.Delete(&bidang).Error

	return bidang, err
}
