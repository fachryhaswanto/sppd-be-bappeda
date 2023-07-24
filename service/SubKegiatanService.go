package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type SubKegiatanService interface {
	FindAll() ([]model.SubKegiatan, error)
	FindById(id int) (model.SubKegiatan, error)
	Create(subKegiatan request.CreateSubKegiatanRequest) (model.SubKegiatan, error)
	Update(id int, subKegiatan request.UpdateSubKegiatanRequest) (model.SubKegiatan, error)
	Delete(id int) (model.SubKegiatan, error)
}

type subKegiatanService struct {
	subKegiatanRepository repository.SubKegiatanRepository
}

func NewSubKegiatanService(repository repository.SubKegiatanRepository) *subKegiatanService {
	return &subKegiatanService{repository}
}

func (s *subKegiatanService) FindAll() ([]model.SubKegiatan, error) {
	var subKegiatans, err = s.subKegiatanRepository.FindAll()

	return subKegiatans, err
}

func (s *subKegiatanService) FindById(id int) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	return subKegiatan, err
}

func (s *subKegiatanService) Create(subKegiatanRequest request.CreateSubKegiatanRequest) (model.SubKegiatan, error) {
	var subKegiatan = model.SubKegiatan{
		KegiatanId:      subKegiatanRequest.KegiatanId,
		KodeSubKegiatan: subKegiatanRequest.KodeSubKegiatan,
		NamaSubKegiatan: subKegiatanRequest.NamaSubKegiatan,
		PejabatId:       subKegiatanRequest.PejabatId,
	}

	newSubKegiatan, err := s.subKegiatanRepository.Create(subKegiatan)

	return newSubKegiatan, err
}

func (s *subKegiatanService) Update(id int, subKegiatanRequest request.UpdateSubKegiatanRequest) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	subKegiatan.KegiatanId = subKegiatanRequest.KegiatanId
	subKegiatan.KodeSubKegiatan = subKegiatanRequest.KodeSubKegiatan
	subKegiatan.NamaSubKegiatan = subKegiatanRequest.NamaSubKegiatan
	subKegiatan.PejabatId = subKegiatanRequest.PejabatId

	updatedSubKegiatan, err := s.subKegiatanRepository.Update(subKegiatan)

	return updatedSubKegiatan, err
}

func (s *subKegiatanService) Delete(id int) (model.SubKegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	deletedSubKegiatan, err := s.subKegiatanRepository.Delete(subKegiatan)

	return deletedSubKegiatan, err
}
