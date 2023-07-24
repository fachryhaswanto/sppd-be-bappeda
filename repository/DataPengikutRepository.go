package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type DataPengikutRepository interface {
	FindAll() ([]model.DataPengikut, error)
	FindById(id int) (model.DataPengikut, error)
	FindBySptId(sptId int) ([]model.DataPengikut, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.DataPengikut, error)
	Create(dataPengikuts []model.DataPengikut) ([]model.DataPengikut, error)
	Delete(dataPengikuts []model.DataPengikut) ([]model.DataPengikut, error)
}

type dataPengikutRepository struct {
	db *gorm.DB
}

func NewDataPengikutRepository(db *gorm.DB) *dataPengikutRepository {
	return &dataPengikutRepository{db}
}

func (r *dataPengikutRepository) FindAll() ([]model.DataPengikut, error) {
	var dataPengikuts []model.DataPengikut

	var err = r.db.Model(&dataPengikuts).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataPengikuts).Error

	return dataPengikuts, err
}

func (r *dataPengikutRepository) FindById(id int) (model.DataPengikut, error) {
	var dataPengikut model.DataPengikut

	var err = r.db.Model(&dataPengikut).Preload("Spt").Preload("Pegawai.Bidang").Take(&dataPengikut, id).Error

	return dataPengikut, err
}

func (r *dataPengikutRepository) FindBySptId(sptId int) ([]model.DataPengikut, error) {
	var dataPengikuts []model.DataPengikut

	var err = r.db.Where("sptId = ?", sptId).Model(&dataPengikuts).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataPengikuts).Error

	return dataPengikuts, err
}

func (r *dataPengikutRepository) FindBySearch(whereClause map[string]interface{}) ([]model.DataPengikut, error) {
	var dataPengikuts []model.DataPengikut

	var err = r.db.Where(whereClause).Model(&dataPengikuts).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataPengikuts).Error

	return dataPengikuts, err
}

func (r *dataPengikutRepository) Create(dataPengikuts []model.DataPengikut) ([]model.DataPengikut, error) {
	var err = r.db.Create(&dataPengikuts).Error

	return dataPengikuts, err
}

func (r *dataPengikutRepository) Delete(dataPengikuts []model.DataPengikut) ([]model.DataPengikut, error) {
	var err error

	for _, dataPengikut := range dataPengikuts {
		err = r.db.Delete(&dataPengikut).Error
	}

	return dataPengikuts, err
}
