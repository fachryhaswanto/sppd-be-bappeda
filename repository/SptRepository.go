package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type SptRepository interface {
	FindAll() ([]model.Spt, error)
	FindById(id int) (model.Spt, error)
	FindAllDitugaskan() ([]model.Spt, error)
	Create(spt model.Spt) (model.Spt, error)
	Update(spt model.Spt) (model.Spt, error)
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

	var err = r.db.Order("id desc").Find(&spts).Error

	return spts, err
}

func (r *sptRepository) FindById(id int) (model.Spt, error) {
	var spt model.Spt

	var err = r.db.Find(&spt, id).Error

	return spt, err
}

func (r *sptRepository) FindAllDitugaskan() ([]model.Spt, error) {
	var ditugaskan []model.Spt

	var err = r.db.Select("id", "tanggal_kembali", "ditugaskan", "status").Where("status = ?", "Belum Kembali").Find(&ditugaskan).Error

	return ditugaskan, err
}

func (r *sptRepository) Create(spt model.Spt) (model.Spt, error) {
	var err = r.db.Create(&spt).Error

	return spt, err
}

func (r *sptRepository) Update(spt model.Spt) (model.Spt, error) {
	var err = r.db.Model(&spt).Updates(model.Spt{
		Template:          spt.Template,
		Nomor_Spt:         spt.Nomor_Spt,
		Tanggal_Spt:       spt.Tanggal_Spt,
		Ditugaskan:        spt.Ditugaskan,
		Jenis_Perjalanan:  spt.Jenis_Perjalanan,
		Keperluan:         spt.Keperluan,
		Alat_Angkutan:     spt.Alat_Angkutan,
		Tempat_Berangkat:  spt.Tempat_Berangkat,
		Tempat_Tujuan:     spt.Tempat_Tujuan,
		Tanggal_Berangkat: spt.Tanggal_Berangkat,
		Tanggal_Kembali:   spt.Tanggal_Kembali,
		Lama_Perjalanan:   spt.Lama_Perjalanan,
		Pejabat_Pemberi:   spt.Pejabat_Pemberi,
		Status:            spt.Status,
		File_Surat_Tugas:  spt.File_Surat_Tugas,
	}).Error

	return spt, err
}

func (r *sptRepository) Delete(spt model.Spt) (model.Spt, error) {
	var err = r.db.Delete(&spt).Error

	return spt, err
}
