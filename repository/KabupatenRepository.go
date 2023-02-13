package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type KabupatenRepository interface {
	FindAll() ([]model.Kabupaten, error)
	FindById(id int) (model.Kabupaten, error)
	Create(kabupaten model.Kabupaten) (model.Kabupaten, error)
	Update(kabupaten model.Kabupaten) (model.Kabupaten, error)
	Delete(kabupaten model.Kabupaten) (model.Kabupaten, error)
}

type kabupatenRepository struct {
	db *gorm.DB
}

func NewKabupatenRepository(db *gorm.DB) *kabupatenRepository {
	return &kabupatenRepository{db}
}

func (r *kabupatenRepository) FindAll() ([]model.Kabupaten, error) {
	var kabupatens []model.Kabupaten

	var err = r.db.Find(&kabupatens).Error

	return kabupatens, err
}

func (r *kabupatenRepository) FindById(id int) (model.Kabupaten, error) {
	var kabupaten model.Kabupaten

	var err = r.db.Find(&kabupaten, id).Error

	return kabupaten, err
}

func (r *kabupatenRepository) Create(kabupaten model.Kabupaten) (model.Kabupaten, error) {
	var err = r.db.Create(&kabupaten).Error

	return kabupaten, err
}

func (r *kabupatenRepository) Update(kabupaten model.Kabupaten) (model.Kabupaten, error) {
	var err = r.db.Model(&kabupaten).Updates(model.Kabupaten{Kode: kabupaten.Kode, Nama: kabupaten.Nama}).Error

	return kabupaten, err
}

func (r *kabupatenRepository) Delete(kabupaten model.Kabupaten) (model.Kabupaten, error) {
	var err = r.db.Delete(&kabupaten).Error

	return kabupaten, err
}
