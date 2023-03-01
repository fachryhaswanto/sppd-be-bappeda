package repository

import (
	"sppd/model"

	"gorm.io/gorm"
)

type ProgramRepository interface {
	FindAll() ([]model.Program, error)
	FindById(id int) (model.Program, error)
	Create(program model.Program) (model.Program, error)
	Update(program model.Program) (model.Program, error)
	Delete(program model.Program) (model.Program, error)
}

type programRepository struct {
	db *gorm.DB
}

func NewProgramRepository(db *gorm.DB) *programRepository {
	return &programRepository{db}
}

func (r *programRepository) FindAll() ([]model.Program, error) {
	var programs []model.Program

	var err = r.db.Find(&programs).Error

	return programs, err
}

func (r *programRepository) FindById(id int) (model.Program, error) {
	var program model.Program

	var err = r.db.Find(&program, id).Error

	return program, err
}

func (r *programRepository) Create(program model.Program) (model.Program, error) {
	var err = r.db.Create(&program).Error

	return program, err
}

func (r *programRepository) Update(program model.Program) (model.Program, error) {
	var err = r.db.Model(&program).Updates(model.Program{
		Kode:       program.Kode,
		Pembebanan: program.Pembebanan,
		Program:    program.Program,
	}).Error

	return program, err
}

func (r *programRepository) Delete(program model.Program) (model.Program, error) {
	var err = r.db.Delete(&program).Error

	return program, err
}
