package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type KegiatanRepository interface {
	FindAll() ([]model.Kegiatan, error)
	FindById(id int) (model.Kegiatan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Kegiatan, error)
	Create(model.Kegiatan) (model.Kegiatan, error)
	Update(model.Kegiatan) (model.Kegiatan, error)
	Delete(model.Kegiatan) (model.Kegiatan, error)
}

type kegiatanRepository struct {
	db *gorm.DB
}

func NewKegiatanRepository(db *gorm.DB) *kegiatanRepository {
	return &kegiatanRepository{db}
}

func (r *kegiatanRepository) FindAll() ([]model.Kegiatan, error) {
	var kegiatans []model.Kegiatan

	var err = r.db.Model(&kegiatans).Preload("Program").Find(&kegiatans).Error

	return kegiatans, err
}

func (r *kegiatanRepository) FindById(id int) (model.Kegiatan, error) {
	var kegiatan model.Kegiatan

	var err = r.db.Model(&kegiatan).Preload("Program").Take(&kegiatan, id).Error

	return kegiatan, err
}

func (r *kegiatanRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Kegiatan, error) {
	var kegiatans []model.Kegiatan

	var err = r.db.Where(whereClause).Model(&kegiatans).Preload("Program").Find(&kegiatans).Error

	return kegiatans, err
}

func (r *kegiatanRepository) Create(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Create(&kegiatan).Error

	return kegiatan, err
}

func (r *kegiatanRepository) Update(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Model(&kegiatan).Updates(model.Kegiatan{
		ProgramId:    kegiatan.ProgramId,
		KodeKegiatan: kegiatan.KodeKegiatan,
		NamaKegiatan: kegiatan.NamaKegiatan,
		Tahun:        kegiatan.Tahun,
	}).Error

	return kegiatan, err
}

func (r *kegiatanRepository) Delete(kegiatan model.Kegiatan) (model.Kegiatan, error) {
	var err = r.db.Delete(&kegiatan).Error

	return kegiatan, err
}
