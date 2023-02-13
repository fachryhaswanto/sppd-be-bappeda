package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type InstansiRepository interface {
	FindAll() ([]model.Instansi, error)
	FindById(id int) (model.Instansi, error)
	Create(model model.Instansi) (model.Instansi, error)
	Update(model model.Instansi) (model.Instansi, error)
	Delete(model model.Instansi) (model.Instansi, error)
}

type instansiRepository struct {
	db *gorm.DB
}

func NewInstansiRepository(db *gorm.DB) *instansiRepository {
	return &instansiRepository{db}
}

func (r *instansiRepository) FindAll() ([]model.Instansi, error) {
	var instansis []model.Instansi

	var err = r.db.Find(&instansis).Error

	return instansis, err
}

func (r *instansiRepository) FindById(id int) (model.Instansi, error) {
	var instansi model.Instansi

	var err = r.db.Find(&instansi, id).Error

	return instansi, err
}

func (r *instansiRepository) Create(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Create(&instansi).Error

	return instansi, err
}

func (r *instansiRepository) Update(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Model(&instansi).Updates(model.Instansi{
		Instansi:      instansi.Instansi,
		Kode_Instansi: instansi.Kode_Instansi,
		Singkatan:     instansi.Singkatan}).Error

	return instansi, err
}

func (r *instansiRepository) Delete(instansi model.Instansi) (model.Instansi, error) {
	var err = r.db.Delete(&instansi).Error

	return instansi, err
}
