package repository

import (
	"sppd/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SptRepository interface {
	FindAll() ([]model.Spt, error)
	FindById(id int) (model.Spt, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Spt, error)
	FindByDate(tahunBerangkat, bulanBerangkat, tahunKembali, bulanKembali string, id int) ([]model.Spt, error)
	CountDataSpt() (int64, error)
	FindAllDitugaskan() ([]model.Spt, error)
	Create(spt model.Spt) (model.Spt, error)
	Update(spt model.Spt) (model.Spt, error)
	UpdateStatusSppd(id int, value int) error
	Delete(spt model.Spt) (model.Spt, error)
}

type sptRepository struct {
	db *gorm.DB
}

func NewSptRepository(db *gorm.DB) *sptRepository {
	return &sptRepository{db}
}

func (r *sptRepository) FindAll() ([]model.Spt, error) {
	var spts []model.Spt

	// var err = r.db.Order("id desc").Find(&spts).Error
	var err = r.db.Model(&spts).Preload("SubKegiatan").Preload("Rekening").Find(&spts).Error

	return spts, err
}

func (r *sptRepository) FindById(id int) (model.Spt, error) {
	var spt model.Spt

	var err = r.db.Model(&spt).Preload(clause.Associations).Preload("SubKegiatan."+clause.Associations).Preload("SubKegiatan.Kegiatan."+clause.Associations).Preload("Rekening").Take(&spt, id).Error

	return spt, err
}

func (r *sptRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Spt, error) {
	var spts []model.Spt

	var err = r.db.Where(whereClause).Model(&spts).Preload("SubKegiatan").Preload("Rekening").Find(&spts).Error

	return spts, err
}

func (r *sptRepository) FindByDate(tahunBerangkat, bulanBerangkat, tahunKembali, bulanKembali string, id int) ([]model.Spt, error) {
	var spts []model.Spt

	var err = r.db.Where("tanggal_berangkat LIKE ? and tanggal_kembali like ? and id != ?", tahunBerangkat+"/"+bulanBerangkat+"%", tahunKembali+"/"+bulanKembali+"%", id).Find(&spts).Error

	return spts, err
}

func (r *sptRepository) FindAllDitugaskan() ([]model.Spt, error) {
	var ditugaskan []model.Spt

	var err = r.db.Select("id", "tanggal_kembali", "ditugaskan", "status").Where("status = ?", "Belum Kembali").Find(&ditugaskan).Error

	return ditugaskan, err
}

func (r *sptRepository) CountDataSpt() (int64, error) {
	var count int64

	var err = r.db.Table("spts").Count(&count).Error

	return count, err
}

func (r *sptRepository) Create(spt model.Spt) (model.Spt, error) {
	var err = r.db.Create(&spt).Error

	return spt, err
}

func (r *sptRepository) Update(spt model.Spt) (model.Spt, error) {
	var err = r.db.Model(&spt).Updates(model.Spt{
		Jenis:             spt.Jenis,
		Template:          spt.Template,
		SubKegiatanId:     spt.SubKegiatanId,
		Nomor_Spt:         spt.Nomor_Spt,
		Tanggal_Spt:       spt.Tanggal_Spt,
		RekeningId:        spt.RekeningId,
		Keperluan:         spt.Keperluan,
		Tanggal_Berangkat: spt.Tanggal_Berangkat,
		Tanggal_Kembali:   spt.Tanggal_Kembali,
		Lama_Perjalanan:   spt.Lama_Perjalanan,
		Pejabat_Pemberi:   spt.Pejabat_Pemberi,
		Status:            spt.Status,
		StatusSppd:        spt.StatusSppd,
		File_Surat_Tugas:  spt.File_Surat_Tugas,
		UserId:            spt.UserId,
	}).Error

	return spt, err
}

func (r *sptRepository) UpdateStatusSppd(id int, value int) error {
	var spt model.Spt

	var err = r.db.Model(&spt).Select("StatusSppd").Where("id = ?", id).Updates(model.Spt{
		StatusSppd: value,
	}).Error

	return err
}

func (r *sptRepository) Delete(spt model.Spt) (model.Spt, error) {
	var err = r.db.Delete(&spt).Error

	return spt, err
}
