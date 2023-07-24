package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type PejabatRepository interface {
	FindAll() ([]model.Pejabat, error)
	FindById(id int) (model.Pejabat, error)
	FindByName(name string) (model.Pejabat, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Pejabat, error)
	Create(pejabat model.Pejabat) (model.Pejabat, error)
	Update(pejabat model.Pejabat) (model.Pejabat, error)
	Delete(pejabat model.Pejabat) (model.Pejabat, error)
}

type pejabatRepository struct {
	db *gorm.DB
}

func NewPejabatRepository(db *gorm.DB) *pejabatRepository {
	return &pejabatRepository{db}
}

func (r *pejabatRepository) FindAll() ([]model.Pejabat, error) {
	var pejabats []model.Pejabat

	var err = r.db.Find(&pejabats).Error

	return pejabats, err
}

func (r *pejabatRepository) FindById(id int) (model.Pejabat, error) {
	var pejabat model.Pejabat

	var err = r.db.Find(&pejabat, id).Error

	return pejabat, err
}

func (r *pejabatRepository) FindByName(name string) (model.Pejabat, error) {
	var pejabat model.Pejabat

	var err = r.db.Where("nama = ?", name).Take(&pejabat).Error

	return pejabat, err
}

func (r *pejabatRepository) FindBySearch(whereClause map[string]interface{}) ([]model.Pejabat, error) {
	var pejabats []model.Pejabat

	var err = r.db.Where(whereClause).Model(&pejabats).Find(&pejabats).Error

	return pejabats, err
}

func (r *pejabatRepository) Create(pejabat model.Pejabat) (model.Pejabat, error) {
	var err = r.db.Create(&pejabat).Error

	return pejabat, err
}

func (r *pejabatRepository) Update(pejabat model.Pejabat) (model.Pejabat, error) {
	var err = r.db.Model(&pejabat).Updates(model.Pejabat{
		Nip:      pejabat.Nip,
		Nama:     pejabat.Nama,
		Pangkat:  pejabat.Pangkat,
		Golongan: pejabat.Golongan,
		Eselon:   pejabat.Eselon,
		Jabatan:  pejabat.Jabatan,
	}).Error

	return pejabat, err
}

func (r *pejabatRepository) Delete(pejabat model.Pejabat) (model.Pejabat, error) {
	var err = r.db.Delete(&pejabat).Error

	return pejabat, err
}
