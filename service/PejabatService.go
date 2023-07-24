package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type PejabatService interface {
	FindAll() ([]model.Pejabat, error)
	FindById(id int) (model.Pejabat, error)
	FindByName(name string) (model.Pejabat, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Pejabat, error)
	Create(pejabat request.CreatePejabatRequest) (model.Pejabat, error)
	Update(id int, pejabat request.UpdatePejabatRequest) (model.Pejabat, error)
	Delete(id int) (model.Pejabat, error)
}

type pejabatService struct {
	pejabatRepository repository.PejabatRepository
}

func NewPejabatService(repository repository.PejabatRepository) *pejabatService {
	return &pejabatService{repository}
}

func (s *pejabatService) FindAll() ([]model.Pejabat, error) {
	var pejabats, err = s.pejabatRepository.FindAll()

	return pejabats, err
}

func (s *pejabatService) FindById(id int) (model.Pejabat, error) {
	var pejabat, err = s.pejabatRepository.FindById(id)

	return pejabat, err
}

func (s *pejabatService) FindByName(name string) (model.Pejabat, error) {
	var pejabat, err = s.pejabatRepository.FindByName(name)

	return pejabat, err
}

func (s *pejabatService) FindBySearch(whereClause map[string]interface{}) ([]model.Pejabat, error) {
	var pejabats, err = s.pejabatRepository.FindBySearch(whereClause)

	return pejabats, err
}

func (s *pejabatService) Create(pejabatRequest request.CreatePejabatRequest) (model.Pejabat, error) {
	var pejabat = model.Pejabat{
		Nip:      pejabatRequest.Nip,
		Nama:     pejabatRequest.Nama,
		Pangkat:  pejabatRequest.Pangkat,
		Golongan: pejabatRequest.Golongan,
		Eselon:   pejabatRequest.Eselon,
		Jabatan:  pejabatRequest.Jabatan,
	}

	newPejabat, err := s.pejabatRepository.Create(pejabat)

	return newPejabat, err
}

func (s *pejabatService) Update(id int, pejabatRequest request.UpdatePejabatRequest) (model.Pejabat, error) {
	var pejabat, err = s.pejabatRepository.FindById(id)

	pejabat.Nip = pejabatRequest.Nip
	pejabat.Nama = pejabatRequest.Nama
	pejabat.Pangkat = pejabatRequest.Pangkat
	pejabat.Golongan = pejabatRequest.Golongan
	pejabat.Eselon = pejabatRequest.Eselon
	pejabat.Jabatan = pejabatRequest.Jabatan

	updatedPejabat, err := s.pejabatRepository.Update(pejabat)

	return updatedPejabat, err
}

func (s *pejabatService) Delete(id int) (model.Pejabat, error) {
	var pejabat, err = s.pejabatRepository.FindById(id)

	deletedPejabat, err := s.pejabatRepository.Delete(pejabat)

	return deletedPejabat, err
}
