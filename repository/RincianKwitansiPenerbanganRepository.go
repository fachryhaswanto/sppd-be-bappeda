package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type RincianKwitansiPenerbanganRepository interface {
	FindAll() ([]model.RincianKwitansiPenerbangan, error)
	FindById(id int) (model.RincianKwitansiPenerbangan, error)
	FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansiPenerbangan, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan) ([]model.RincianKwitansiPenerbangan, error)
	Update(rincianKwitansiPenerbangan model.RincianKwitansiPenerbangan) (model.RincianKwitansiPenerbangan, error)
	Delete(rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan) ([]model.RincianKwitansiPenerbangan, error)
}

type rincianKwitansiPenerbanganRepository struct {
	db *gorm.DB
}

func NewRincianKwitansiPenerbanganRepository(db *gorm.DB) *rincianKwitansiPenerbanganRepository {
	return &rincianKwitansiPenerbanganRepository{db}
}

func (r *rincianKwitansiPenerbanganRepository) FindAll() ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan

	var err = r.db.Model(&rincianKwitansiPenerbangans).Preload("Kwitansi").Find(&rincianKwitansiPenerbangans).Error

	return rincianKwitansiPenerbangans, err
}

func (r *rincianKwitansiPenerbanganRepository) FindById(id int) (model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangan model.RincianKwitansiPenerbangan

	var err = r.db.Model(&rincianKwitansiPenerbangan).Preload("Kwitansi").Take(&rincianKwitansiPenerbangan, id).Error

	return rincianKwitansiPenerbangan, err
}

func (r *rincianKwitansiPenerbanganRepository) FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan

	var err = r.db.Where("kwitansiId = ?", kwitansiId).Model(&rincianKwitansiPenerbangans).Preload("Kwitansi").Find(&rincianKwitansiPenerbangans).Error

	return rincianKwitansiPenerbangans, err
}

func (r *rincianKwitansiPenerbanganRepository) FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan

	var err = r.db.Where(whereClause).Model(&rincianKwitansiPenerbangans).Preload("Kwitansi").Find(&rincianKwitansiPenerbangans).Error

	return rincianKwitansiPenerbangans, err
}

func (r *rincianKwitansiPenerbanganRepository) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count int64
	var err error
	var _, ok = whereClause["all"]

	if ok {
		err = r.db.Table("rincian_kwitansi_penerbangans").Count(&count).Error
	} else {
		err = r.db.Where(whereClause).Table("rincian_kwitansi_penerbangans").Count(&count).Error
	}

	return count, err
}

func (r *rincianKwitansiPenerbanganRepository) Create(rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan) ([]model.RincianKwitansiPenerbangan, error) {
	var err = r.db.Create(&rincianKwitansiPenerbangans).Error

	return rincianKwitansiPenerbangans, err
}

func (r *rincianKwitansiPenerbanganRepository) Update(rincianKwitansiPenerbangan model.RincianKwitansiPenerbangan) (model.RincianKwitansiPenerbangan, error) {
	var err = r.db.Model(&rincianKwitansiPenerbangan).Updates(model.RincianKwitansiPenerbangan{
		Id:           rincianKwitansiPenerbangan.Id,
		KwitansiId:   rincianKwitansiPenerbangan.KwitansiId,
		NamaMaskapai: rincianKwitansiPenerbangan.NamaMaskapai,
		NomorTiket:   rincianKwitansiPenerbangan.NomorTiket,
		KodeBooking:  rincianKwitansiPenerbangan.KodeBooking,
	}).Error

	return rincianKwitansiPenerbangan, err
}

func (r *rincianKwitansiPenerbanganRepository) Delete(rincianKwitansiPenerbangans []model.RincianKwitansiPenerbangan) ([]model.RincianKwitansiPenerbangan, error) {
	var err error

	for _, rincianKwitansiPenerbangan := range rincianKwitansiPenerbangans {
		err = r.db.Delete(&rincianKwitansiPenerbangan).Error
	}

	return rincianKwitansiPenerbangans, err
}
