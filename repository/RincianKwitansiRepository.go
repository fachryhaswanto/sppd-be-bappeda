package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type RincianKwitansiRepository interface {
	FindAll() ([]model.RincianKwitansi, error)
	FindById(id int) (model.RincianKwitansi, error)
	FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansi, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(rincianKwitansi []model.RincianKwitansi) ([]model.RincianKwitansi, error)
	Update(rincianKwitansi model.RincianKwitansi) (model.RincianKwitansi, error)
	Delete(rincianKwitansis []model.RincianKwitansi) ([]model.RincianKwitansi, error)
}

type rincianKwitansiRepository struct {
	db *gorm.DB
}

func NewRincianKwitansiRepository(db *gorm.DB) *rincianKwitansiRepository {
	return &rincianKwitansiRepository{db}
}

func (r *rincianKwitansiRepository) FindAll() ([]model.RincianKwitansi, error) {
	var rincianKwitansis []model.RincianKwitansi

	var err = r.db.Model(&rincianKwitansis).Preload("Kwitansi").Find(&rincianKwitansis).Error

	return rincianKwitansis, err
}

func (r *rincianKwitansiRepository) FindById(id int) (model.RincianKwitansi, error) {
	var rincianKwitansi model.RincianKwitansi

	var err = r.db.Model(&rincianKwitansi).Preload("Kwitansi").Take(&rincianKwitansi, id).Error

	return rincianKwitansi, err
}

func (r *rincianKwitansiRepository) FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansi, error) {
	var rincianKwitansis []model.RincianKwitansi

	var err = r.db.Where("kwitansiId = ?", kwitansiId).Model(&rincianKwitansis).Preload("Kwitansi").Find(&rincianKwitansis).Error

	return rincianKwitansis, err
}

func (r *rincianKwitansiRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansi, error) {
	var rincianKwitansis []model.RincianKwitansi

	var err = r.db.Where(whereClause).Model(&rincianKwitansis).Preload("Kwitansi").Find(&rincianKwitansis).Error

	return rincianKwitansis, err
}

func (r *rincianKwitansiRepository) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count int64
	var err error
	var _, ok = whereClause["all"]

	if ok {
		err = r.db.Table("rincian_kwitansis").Count(&count).Error
	} else {
		err = r.db.Where(whereClause).Table("rincian_kwitansis").Count(&count).Error
	}

	return count, err
}

func (r *rincianKwitansiRepository) Create(rincianKwitansi []model.RincianKwitansi) ([]model.RincianKwitansi, error) {
	var err = r.db.Create(&rincianKwitansi).Error

	return rincianKwitansi, err
}

func (r *rincianKwitansiRepository) Update(rincianKwitansi model.RincianKwitansi) (model.RincianKwitansi, error) {
	var err = r.db.Model(&rincianKwitansi).Updates(model.RincianKwitansi{
		Id:          rincianKwitansi.Id,
		KwitansiId:  rincianKwitansi.KwitansiId,
		NamaRincian: rincianKwitansi.NamaRincian,
		JumlahBayar: rincianKwitansi.JumlahBayar,
		Banyaknya:   rincianKwitansi.Banyaknya,
		HasilBayar:  rincianKwitansi.HasilBayar,
	}).Error

	return rincianKwitansi, err
}

func (r *rincianKwitansiRepository) Delete(rincianKwitansis []model.RincianKwitansi) ([]model.RincianKwitansi, error) {
	var err error

	for _, rincianKwitansi := range rincianKwitansis {
		err = r.db.Delete(&rincianKwitansi).Error
	}

	return rincianKwitansis, err
}
