package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type RekeningRepository interface {
	FindAll() ([]model.Rekening, error)
	FindById(id int) (model.Rekening, error)
	Create(rekening model.Rekening) (model.Rekening, error)
	Update(rekening model.Rekening) (model.Rekening, error)
	Delete(rekening model.Rekening) (model.Rekening, error)
}

type rekeningRepository struct {
	db *gorm.DB
}

func NewRekeningRepository(db *gorm.DB) *rekeningRepository {
	return &rekeningRepository{db}
}

func (r *rekeningRepository) FindAll() ([]model.Rekening, error) {
	var rekenings []model.Rekening

	var err = r.db.Find(&rekenings).Error

	return rekenings, err
}

func (r *rekeningRepository) FindById(id int) (model.Rekening, error) {
	var rekening model.Rekening

	var err = r.db.Take(&rekening, id).Error

	return rekening, err
}

func (r *rekeningRepository) Create(rekening model.Rekening) (model.Rekening, error) {
	var err = r.db.Create(&rekening).Error

	return rekening, err
}

func (r *rekeningRepository) Update(rekening model.Rekening) (model.Rekening, error) {
	var err = r.db.Model(&rekening).Updates(model.Rekening{
		KodeRekening: rekening.KodeRekening,
		NamaRekening: rekening.NamaRekening,
	}).Error

	return rekening, err
}

func (r *rekeningRepository) Delete(rekening model.Rekening) (model.Rekening, error) {
	var err = r.db.Delete(&rekening).Error

	return rekening, err
}
