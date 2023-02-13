package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type KabupatenService interface {
	FindAll() ([]model.Kabupaten, error)
	FindById(id int) (model.Kabupaten, error)
	Create(kabupaten request.CreateKabupatenRequest) (model.Kabupaten, error)
	Update(id int, kabupaten request.UpdateKabupatenRequest) (model.Kabupaten, error)
	Delete(id int) (model.Kabupaten, error)
}

type kabupatenService struct {
	kabupatenRepository repository.KabupatenRepository
}

func NewKabupatenService(repository repository.KabupatenRepository) *kabupatenService {
	return &kabupatenService{repository}
}

func (s *kabupatenService) FindAll() ([]model.Kabupaten, error) {
	var members, err = s.kabupatenRepository.FindAll()

	return members, err
}

func (s *kabupatenService) FindById(id int) (model.Kabupaten, error) {
	var member, err = s.kabupatenRepository.FindById(id)

	return member, err
}

func (s *kabupatenService) Create(kabupatenRequest request.CreateKabupatenRequest) (model.Kabupaten, error) {
	var kabupaten = model.Kabupaten{
		Kode: kabupatenRequest.Kode,
		Nama: kabupatenRequest.Nama,
	}

	var newKabupaten, err = s.kabupatenRepository.Create(kabupaten)

	return newKabupaten, err
}

func (s *kabupatenService) Update(id int, kabupatenRequest request.UpdateKabupatenRequest) (model.Kabupaten, error) {
	var kabupaten, err = s.kabupatenRepository.FindById(id)

	kabupaten.Kode = kabupatenRequest.Kode
	kabupaten.Nama = kabupatenRequest.Nama

	updatedKabupaten, err := s.kabupatenRepository.Update(kabupaten)

	return updatedKabupaten, err
}

func (s *kabupatenService) Delete(id int) (model.Kabupaten, error) {
	var member, err = s.kabupatenRepository.FindById(id)

	deletedMember, err := s.kabupatenRepository.Delete(member)

	return deletedMember, err
}
