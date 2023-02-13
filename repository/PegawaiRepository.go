package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type PegawaiRepository interface {
	FindAll() ([]model.Pegawai, error)
	FindById(id int) (model.Pegawai, error)
	FindByName(name string) (model.Pegawai, error)
	Create(pegawai model.Pegawai) (model.Pegawai, error)
	Update(pegawai model.Pegawai) (model.Pegawai, error)
	Delete(pegawai model.Pegawai) (model.Pegawai, error)
}

type pegawaiRepository struct {
	db *gorm.DB
}

func NewPegawaiRepository(db *gorm.DB) *pegawaiRepository {
	return &pegawaiRepository{db}
}

func (r *pegawaiRepository) FindAll() ([]model.Pegawai, error) {
	var pegawais []model.Pegawai

	var err = r.db.Find(&pegawais).Error

	return pegawais, err
}

func (r *pegawaiRepository) FindById(id int) (model.Pegawai, error) {
	var pegawai model.Pegawai

	var err = r.db.Find(&pegawai, id).Error

	return pegawai, err
}

func (r *pegawaiRepository) FindByName(name string) (model.Pegawai, error) {
	var pegawai model.Pegawai

	var err = r.db.Where("nama = ?", name).Take(&pegawai).Error

	return pegawai, err
}

func (r *pegawaiRepository) Create(pegawai model.Pegawai) (model.Pegawai, error) {
	var err = r.db.Create(&pegawai).Error

	return pegawai, err
}

func (r *pegawaiRepository) Update(pegawai model.Pegawai) (model.Pegawai, error) {
	var err = r.db.Model(&pegawai).Updates(model.Pegawai{
		Nip:           pegawai.Nip,
		Nama:          pegawai.Nama,
		Jenis_Kelamin: pegawai.Jenis_Kelamin,
		Status:        pegawai.Status,
		Tempat_Lahir:  pegawai.Tempat_Lahir,
		Tanggal_Lahir: pegawai.Tanggal_Lahir,
		Instansi:      pegawai.Instansi,
		Bidang:        pegawai.Bidang,
		Golongan:      pegawai.Golongan,
		Eselon:        pegawai.Eselon,
		Pangkat:       pegawai.Pangkat,
		Jabatan:       pegawai.Jabatan,
	}).Error

	return pegawai, err
}

func (r *pegawaiRepository) Delete(pegawai model.Pegawai) (model.Pegawai, error) {
	var err = r.db.Delete(&pegawai).Error

	return pegawai, err
}
