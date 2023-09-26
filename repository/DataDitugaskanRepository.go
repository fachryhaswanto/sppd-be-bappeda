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
	FindPengikut(pegawaiId, sptId int) ([]model.DataDitugaskan, error)
	CountDataBySptId(sptId int) (int64, error)
	CountDataByPegawaiId(dataPegawai []model.Pegawai) []int64
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(dataDitugaskan []model.DataDitugaskan) ([]model.DataDitugaskan, error)
	UpdateStatusSppd(sptId int, pegawaiId int, value int) error
	UpdateStatusKwitansi(sptId int, pegawaiId int, value int) error
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

func (r *dataDitugaskanRepository) FindPengikut(pegawaiId, sptId int) ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan

	var err = r.db.Where("pegawaiId != ? and sptId = ?", pegawaiId, sptId).Model(&dataDitugaskans).Preload("Spt").Preload("Pegawai.Bidang").Find(&dataDitugaskans).Error

	return dataDitugaskans, err
}

func (r *dataDitugaskanRepository) CountDataBySptId(sptId int) (int64, error) {
	var dataDitugaskan model.DataDitugaskan
	var count int64

	var err = r.db.Model(&dataDitugaskan).Where("sptId = ?", sptId).Count(&count).Error

	return count, err
}

func (r *dataDitugaskanRepository) CountDataByPegawaiId(dataPegawai []model.Pegawai) []int64 {
	var arrJumlahPerjalanan []int64
	var jumlahPerjalanan int64

	for _, data := range dataPegawai {
		r.db.Table("data_ditugaskans").Where("pegawaiId = ?", data.Id).Count(&jumlahPerjalanan)

		arrJumlahPerjalanan = append(arrJumlahPerjalanan, jumlahPerjalanan)
	}

	return arrJumlahPerjalanan
}

func (r *dataDitugaskanRepository) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var dataDitugaskan model.DataDitugaskan
	var count int64

	var err = r.db.Model(&dataDitugaskan).Where(whereClause).Count(&count).Error

	return count, err
}

func (r *dataDitugaskanRepository) Create(dataDitugaskan []model.DataDitugaskan) ([]model.DataDitugaskan, error) {
	var err = r.db.Create(&dataDitugaskan).Error

	return dataDitugaskan, err
}

func (r *dataDitugaskanRepository) UpdateStatusSppd(sptId int, pegawaiId int, value int) error {
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

func (r *dataDitugaskanRepository) UpdateStatusKwitansi(sptId int, pegawaiId int, value int) error {
	var dataDitugaskan model.DataDitugaskan
	var err error

	if pegawaiId != 0 {
		err = r.db.Model(&dataDitugaskan).Select("StatusKwitansi").Where("sptId = ? and pegawaiId = ?", sptId, pegawaiId).Updates(model.DataDitugaskan{
			StatusKwitansi: value,
		}).Error
	} else {
		err = r.db.Model(&dataDitugaskan).Select("StatusKwitansi").Where("sptId = ?", sptId).Updates(model.DataDitugaskan{
			StatusKwitansi: value,
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
