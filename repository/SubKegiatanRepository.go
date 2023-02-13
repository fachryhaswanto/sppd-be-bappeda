package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type SubKegiatanRepository interface {
	FindAll() ([]model.Sub_Kegiatan, error)
	FindById(id int) (model.Sub_Kegiatan, error)
	Create(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
	Update(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
	Delete(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error)
}

type subKegiatanRepository struct {
	db *gorm.DB
}

func NewSubKegiatanRepository(db *gorm.DB) *subKegiatanRepository {
	return &subKegiatanRepository{db}
}

func (r *subKegiatanRepository) FindAll() ([]model.Sub_Kegiatan, error) {
	var subKegiatans []model.Sub_Kegiatan

	var err = r.db.Find(&subKegiatans).Error

	return subKegiatans, err
}

func (r *subKegiatanRepository) FindById(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan model.Sub_Kegiatan

	var err = r.db.Find(&subKegiatan, id).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Create(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Create(&subKegiatan).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Update(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Model(&subKegiatan).Updates(model.Sub_Kegiatan{
		Tahun:        subKegiatan.Tahun,
		Nama_Program: subKegiatan.Nama_Program,
		Kode:         subKegiatan.Kode,
		Kegiatan:     subKegiatan.Kegiatan,
	}).Error

	return subKegiatan, err
}

func (r *subKegiatanRepository) Delete(subKegiatan model.Sub_Kegiatan) (model.Sub_Kegiatan, error) {
	var err = r.db.Delete(&subKegiatan).Error

	return subKegiatan, err
}
