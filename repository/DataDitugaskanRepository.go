package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type DataDitugaskanRepository interface {
	FindAll() ([]model.DataDitugaskan, error)
	FindById(id int) (model.DataDitugaskan, error)
	FindBySptId(sptId int) ([]model.DataDitugaskan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.DataDitugaskan, error)
	FindByPegawaiId(pegawaiId int) ([]model.DataDitugaskan, error)
	CountDataBySptId(sptId int) (int64, error)
	Create(dataDitugaskan []model.DataDitugaskan) ([]model.DataDitugaskan, error)
	UpdateStatus(sptId int, pegawaiId int, value int) error
	Delete(dataDitugaskans []model.DataDitugaskan) ([]model.DataDitugaskan, error)
}

type dataDitugaskanRepository struct {
	db *gorm.DB
}

func NewDataDitugaskanRepository(db *gorm.DB) *dataDitugaskanRepository {
	return &dataDitugaskanRepository{db}
}

func (r *dataDitugaskanRepository) FindAll() ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan

	var err = r.db.Model(&dataDitugaskans).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataDitugaskans).Error

	return dataDitugaskans, err
}

func (r *dataDitugaskanRepository) FindById(id int) (model.DataDitugaskan, error) {
	var dataDitugaskan model.DataDitugaskan

	var err = r.db.Model(&dataDitugaskan).Preload("Spt").Preload("Pegawai.Bidang").Take(&dataDitugaskan, id).Error

	return dataDitugaskan, err
}

func (r *dataDitugaskanRepository) FindBySptId(sptId int) ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan

	var err = r.db.Where("sptId = ?", sptId).Model(&dataDitugaskans).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataDitugaskans).Error

	return dataDitugaskans, err
}

func (r *dataDitugaskanRepository) FindBySearch(whereClause map[string]interface{}) ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan

	var err = r.db.Where(whereClause).Model(&dataDitugaskans).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataDitugaskans).Error

	return dataDitugaskans, err
}

func (r *dataDitugaskanRepository) FindByPegawaiId(pegawaiId int) ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan

	var err = r.db.Where("pegawaiId = ?", pegawaiId).Model(&dataDitugaskans).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataDitugaskans).Error

	return dataDitugaskans, err
}

func (r *dataDitugaskanRepository) CountDataBySptId(sptId int) (int64, error) {
	var dataDitugaskan model.DataDitugaskan
	var count int64

	var err = r.db.Model(&dataDitugaskan).Where("sptId = ?", sptId).Count(&count).Error

	return count, err
}

func (r *dataDitugaskanRepository) Create(dataDitugaskan []model.DataDitugaskan) ([]model.DataDitugaskan, error) {
	var err = r.db.Create(&dataDitugaskan).Error

	return dataDitugaskan, err
}

func (r *dataDitugaskanRepository) UpdateStatus(sptId int, pegawaiId int, value int) error {
	var dataDitugaskan model.DataDitugaskan
	var err error

	if pegawaiId != 0 {
		err = r.db.Model(&dataDitugaskan).Select("StatusSppd").Where("sptId = ? and pegawaiId = ?", sptId, pegawaiId).Updates(model.DataDitugaskan{
			StatusSppd: value,
		}).Error
	} else {
		err = r.db.Model(&dataDitugaskan).Select("StatusSppd").Where("sptId = ?", sptId).Updates(model.DataDitugaskan{
			StatusSppd: value,
		}).Error
	}

	return err
}

func (r *dataDitugaskanRepository) Delete(dataDitugaskans []model.DataDitugaskan) ([]model.DataDitugaskan, error) {
	var err error

	for _, dataDitugaskan := range dataDitugaskans {
		err = r.db.Delete(&dataDitugaskan).Error
	}

	return dataDitugaskans, err
}
