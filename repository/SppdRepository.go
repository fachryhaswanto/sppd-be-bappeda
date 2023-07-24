package repository

import (
	"sppd/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SppdRepository interface {
	FindAll() ([]model.Sppd, error)
	FindById(id int) (model.Sppd, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Sppd, error)
	Create(sppd model.Sppd) (model.Sppd, error)
	CountDataBySptId(sptId int) (int64, error)
	Update(sppd model.Sppd) (model.Sppd, error)
	UpdateBySptId(sptId int, sppd model.Sppd) (model.Sppd, error)
	Delete(sppd model.Sppd) (model.Sppd, error)
	DeleteBySptId(sptId int) error
}

type sppdRepository struct {
	db *gorm.DB
}

func NewSppdRepository(db *gorm.DB) *sppdRepository {
	return &sppdRepository{db}
}

func (r *sppdRepository) FindAll() ([]model.Sppd, error) {
	var sppds []model.Sppd

	var err = r.db.Model(&sppds).Preload(clause.Associations).Preload("Pegawai").Preload("Pejabat").Preload("Spt." + clause.Associations).Preload("Spt.SubKegiatan." + clause.Associations).Preload("Spt.SubKegiatan.Pejabat." + clause.Associations).Preload("Spt.Rekening." + clause.Associations).Find(&sppds).Error

	return sppds, err
}

// func (r *sppdRepository) FindJoin() ([]model.JoinSppdSpt, error) {

// 	var join []model.JoinSppdSpt
// 	var sppd model.Sppd

// 	var err = r.db.Model(&sppd).Select("sppds.id, sppds.template_sppd, spts.ditugaskan, sppds.nomor_sppd, sppds.tanggal_sppd, sppds.tingkat_biaya, sppds.instansi, sppds.tanda_tangan, spts.nomor_spt, spts.tanggal_spt, spts.jenis_perjalanan, spts.keperluan, spts.alat_angkutan, spts.tempat_berangkat, spts.tempat_tujuan, spts.tanggal_berangkat, spts.tanggal_kembali, spts.lama_perjalanan, spts.pejabat_pemberi, spts.status").Joins("left join spts on sppds.idspt = spts.id").Order("id DESC").Scan(&join).Error

// 	return join, err
// }

func (r *sppdRepository) FindById(id int) (model.Sppd, error) {
	var sppd model.Sppd

	var err = r.db.Model(&sppd).Preload(clause.Associations).Preload("Pegawai").Preload("Pejabat").Preload("Spt."+clause.Associations).Preload("Spt.SubKegiatan."+clause.Associations).Preload("Spt.SubKegiatan.Kegiatan."+clause.Associations).Preload("Spt.SubKegiatan.Kegiatan.Program."+clause.Associations).Preload("Spt.SubKegiatan.Pejabat."+clause.Associations).Preload("Spt.Rekening."+clause.Associations).Take(&sppd, id).Error

	return sppd, err
}

func (r *sppdRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Sppd, error) {
	var sppds []model.Sppd

	var err = r.db.Where(whereClause).Model(&sppds).Preload(clause.Associations).Preload("Pegawai").Preload("Pejabat").Preload("Spt." + clause.Associations).Preload("Spt.SubKegiatan." + clause.Associations).Preload("Spt.SubKegiatan.Pejabat." + clause.Associations).Preload("Spt.Rekening." + clause.Associations).Find(&sppds).Error

	return sppds, err
}

func (r *sppdRepository) CountDataBySptId(sptId int) (int64, error) {
	var sppd model.Sppd
	var count int64

	var err = r.db.Model(&sppd).Where("sptId = ?", sptId).Count(&count).Error

	return count, err
}

func (r *sppdRepository) Create(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Create(&sppd).Error

	return sppd, err
}

func (r *sppdRepository) Update(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Model(&sppd).Updates(model.Sppd{
		Template_Sppd:    sppd.Template_Sppd,
		Jenis:            sppd.Jenis,
		Nomor_Sppd:       sppd.Nomor_Sppd,
		PegawaiId:        sppd.PegawaiId,
		Tanggal_Sppd:     sppd.Tanggal_Sppd,
		Tempat_Berangkat: sppd.Tempat_Berangkat,
		Tempat_Tujuan:    sppd.Tempat_Tujuan,
		Alat_Angkutan:    sppd.Alat_Angkutan,
		Instansi:         sppd.Instansi,
		PejabatId:        sppd.PejabatId,
		SptId:            sppd.SptId,
		UserId:           sppd.UserId,
	}).Error

	return sppd, err
}

func (r *sppdRepository) UpdateBySptId(sptId int, sppd model.Sppd) (model.Sppd, error) {

	var err = r.db.Model(&sppd).Where("sptId = ?", sptId).Updates(model.Sppd{
		Template_Sppd:    sppd.Template_Sppd,
		Jenis:            sppd.Jenis,
		Nomor_Sppd:       sppd.Nomor_Sppd,
		PegawaiId:        sppd.PegawaiId,
		Tanggal_Sppd:     sppd.Tanggal_Sppd,
		Tempat_Berangkat: sppd.Tempat_Berangkat,
		Tempat_Tujuan:    sppd.Tempat_Tujuan,
		Alat_Angkutan:    sppd.Alat_Angkutan,
		Instansi:         sppd.Instansi,
		PejabatId:        sppd.PejabatId,
		SptId:            sppd.SptId,
	}).Error

	return sppd, err
}

func (r *sppdRepository) Delete(sppd model.Sppd) (model.Sppd, error) {
	var err = r.db.Delete(&sppd).Error

	return sppd, err
}

func (r *sppdRepository) DeleteBySptId(sptId int) error {
	var sppd model.Sppd

	var err = r.db.Where("sptId = ?", sptId).Delete(&sppd).Error

	return err
}
