package repository

import (
	"database/sql"
	"sppd/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KwitansiRepository interface {
	FindAll() ([]model.Kwitansi, error)
	FindById(id int) (model.Kwitansi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Kwitansi, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	SumTotalBayar(whereClause map[string]interface{}) int64
	Create(kwitansi model.Kwitansi) (model.Kwitansi, error)
	Update(kwitansi model.Kwitansi) (model.Kwitansi, error)
	UpdateTotalBayar(kwitansi model.Kwitansi) (model.Kwitansi, error)
	Delete(kwitansi model.Kwitansi) (model.Kwitansi, error)
}

type kwitansiRepository struct {
	db *gorm.DB
}

func NewKwitansiRepository(db *gorm.DB) *kwitansiRepository {
	return &kwitansiRepository{db}
}

func (r *kwitansiRepository) FindAll() ([]model.Kwitansi, error) {
	var kwitansis []model.Kwitansi

	var err = r.db.Model(&kwitansis).Preload(clause.Associations).Preload("Sppd").Preload("Sppd." + clause.Associations).Preload("Sppd.Spt." + clause.Associations).Preload("Sppd.Spt.SubKegiatan." + clause.Associations).Find(&kwitansis).Error

	return kwitansis, err
}

func (r *kwitansiRepository) FindById(id int) (model.Kwitansi, error) {
	var kwitansi model.Kwitansi

	var err = r.db.Model(&kwitansi).Preload(clause.Associations).Preload("Sppd").Preload("Sppd."+clause.Associations).Preload("Sppd.Spt."+clause.Associations).Preload("Sppd.Spt.SubKegiatan."+clause.Associations).Take(&kwitansi, id).Error

	return kwitansi, err
}

func (r *kwitansiRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Kwitansi, error) {
	var kwitansis []model.Kwitansi

	var err = r.db.Where(whereClause).Model(&kwitansis).Preload(clause.Associations).Preload("Sppd").Preload("Sppd." + clause.Associations).Preload("Sppd.Spt." + clause.Associations).Preload("Sppd.Spt.SubKegiatan." + clause.Associations).Find(&kwitansis).Error

	return kwitansis, err
}

func (r *kwitansiRepository) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count int64
	var err error
	var _, ok = whereClause["all"]

	if ok {
		err = r.db.Table("kwitansis").Count(&count).Error
	} else {
		err = r.db.Where(whereClause).Table("kwitansis").Count(&count).Error
	}

	return count, err
}

func (r *kwitansiRepository) SumTotalBayar(whereClause map[string]interface{}) int64 {
	var realisasi int64
	var _, ok = whereClause["all"]
	var rows *sql.Rows

	if ok {
		rows, _ = r.db.Table("kwitansis").Select("sum(totalBayar) as realisasi").Rows()
	} else {
		rows, _ = r.db.Table("kwitansis").Where(whereClause).Select("sum(totalBayar) as realisasi").Rows()
	}

	for rows.Next() {
		rows.Scan(&realisasi)
	}

	return realisasi
}

func (r *kwitansiRepository) Create(kwitansi model.Kwitansi) (model.Kwitansi, error) {
	var err = r.db.Create(&kwitansi).Error

	return kwitansi, err
}

func (r *kwitansiRepository) Update(kwitansi model.Kwitansi) (model.Kwitansi, error) {
	var err = r.db.Model(&kwitansi).Select("totalBayar").Updates(model.Kwitansi{
		SppdId:        kwitansi.SppdId,
		PegawaiId:     kwitansi.PegawaiId,
		NomorKwitansi: kwitansi.NomorKwitansi,
		TanggalBayar:  kwitansi.TanggalBayar,
		Keperluan:     kwitansi.Keperluan,
		TotalBayar:    kwitansi.TotalBayar,
		Tahun:         kwitansi.Tahun,
		UserId:        kwitansi.UserId,
	}).Error

	return kwitansi, err
}

func (r *kwitansiRepository) UpdateTotalBayar(kwitansi model.Kwitansi) (model.Kwitansi, error) {
	var err = r.db.Model(&kwitansi).Select("totalBayar").Updates(model.Kwitansi{
		TotalBayar: kwitansi.TotalBayar,
	}).Error

	return kwitansi, err
}

func (r *kwitansiRepository) Delete(kwitansi model.Kwitansi) (model.Kwitansi, error) {
	var err = r.db.Delete(&kwitansi).Error

	return kwitansi, err
}
