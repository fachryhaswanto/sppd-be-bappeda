package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type SppdRepository interface {
	FindAll() ([]model.Sppd, error)
	FindJoin() ([]model.JoinSppdSpt, error)
	FindById(id int) (model.Sppd, error)
	Create(sppd model.Sppd) (model.Sppd, error)
	Update(sppd model.Sppd) (model.Sppd, error)
	Delete(sppd model.Sppd) (model.Sppd, error)
}

type sppdRepository struct {
	db *gorm.DB
}

func NewSppdRepository(db *gorm.DB) *sppdRepository {
	return &sppdRepository{db}
}

func (r *sppdRepository) FindAll() ([]model.Sppd, error) {
	var sppds []model.Sppd

	var err = r.db.Find(&sppds).Error

	return sppds, err
}
func (r *sppdRepository) FindJoin() ([]model.JoinSppdSpt, error) {

	var join []model.JoinSppdSpt
	var sppd model.Sppd

	var err = r.db.Model(&sppd).Select("sppds.id, sppds.template_sppd, spts.ditugaskan, sppds.nomor_sppd, sppds.tanggal_sppd, sppds.tingkat_biaya, sppds.instansi, sppds.tanda_tangan, spts.nomor_spt, spts.tanggal_spt, spts.jenis_perjalanan, spts.keperluan, spts.alat_angkutan, spts.tempat_berangkat, spts.tempat_tujuan, spts.tanggal_berangkat, spts.tanggal_kembali, spts.lama_perjalanan, spts.pejabat_pemberi, spts.status").Joins("left join spts on sppds.idspt = spts.id").Order("id DESC").Scan(&join).Error

	return join, err
}

func (r *sppdRepository) FindById(id int) (model.Sppd, error) {
	var sppd model.Sppd

	var err = r.db.Find(&sppd, id).Error

	return sppd, err
}

func (r *sppdRepository) Create(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Create(&sppd).Error

	return sppd, err
}

func (r *sppdRepository) Update(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Model(&sppd).Updates(model.Sppd{
		Template_Sppd: sppd.Template_Sppd,
		Nomor_Sppd:    sppd.Nomor_Sppd,
		Tanggal_Sppd:  sppd.Tanggal_Sppd,
		Tingkat_Biaya: sppd.Tingkat_Biaya,
		Instansi:      sppd.Instansi,
		Tanda_Tangan:  sppd.Tanda_Tangan,
		Idspt:         sppd.Idspt,
	}).Error

	return sppd, err
}

func (r *sppdRepository) Delete(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Delete(&sppd).Error

	return sppd, err
}
