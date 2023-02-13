package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type ProgramService interface {
	FindAll() ([]model.Program, error)
	FindById(id int) (model.Program, error)
	Create(program request.CreateProgramRequest) (model.Program, error)
	Update(id int, program request.UpdateProgramRequest) (model.Program, error)
	Delete(id int) (model.Program, error)
}

type programService struct {
	programRepository repository.ProgramRepository
}

func NewProgramService(repository repository.ProgramRepository) *programService {
	return &programService{repository}
}

func (s *programService) FindAll() ([]model.Program, error) {
	var programs, err = s.programRepository.FindAll()

	return programs, err
}

func (s *programService) FindById(id int) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	return program, err
}

func (s *programService) Create(programRequest request.CreateProgramRequest) (model.Program, error) {
	var program = model.Program{
		Tahun:      programRequest.Tahun,
		Kode:       programRequest.Kode,
		Pembebanan: programRequest.Pembebanan,
		Program:    programRequest.Program,
	}

	var newProgram, err = s.programRepository.Create(program)

	return newProgram, err
}

func (s *programService) Update(id int, programRequest request.UpdateProgramRequest) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	program.Tahun = programRequest.Tahun
	program.Kode = programRequest.Kode
	program.Pembebanan = programRequest.Pembebanan
	program.Program = programRequest.Program

	updatedPegawai, err := s.programRepository.Update(program)

	return updatedPegawai, err
}

func (s *programService) Delete(id int) (model.Program, error) {
	var program, err = s.programRepository.FindById(id)

	deletedProgram, err := s.programRepository.Delete(program)

	return deletedProgram, err
}
