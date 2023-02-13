package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type SubKegiatanService interface {
	FindAll() ([]model.Sub_Kegiatan, error)
	FindById(id int) (model.Sub_Kegiatan, error)
	Create(subKegiatan request.CreateSubKegiatanRequest) (model.Sub_Kegiatan, error)
	Update(id int, subKegiatan request.UpdateSubKegiatanRequest) (model.Sub_Kegiatan, error)
	Delete(id int) (model.Sub_Kegiatan, error)
}

type subKegiatanService struct {
	subKegiatanRepository repository.SubKegiatanRepository
}

func NewSubKegiatanService(repository repository.SubKegiatanRepository) *subKegiatanService {
	return &subKegiatanService{repository}
}

func (s *subKegiatanService) FindAll() ([]model.Sub_Kegiatan, error) {
	var subKegiatans, err = s.subKegiatanRepository.FindAll()

	return subKegiatans, err
}

func (s *subKegiatanService) FindById(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	return subKegiatan, err
}

func (s *subKegiatanService) Create(subKegiatanRequest request.CreateSubKegiatanRequest) (model.Sub_Kegiatan, error) {
	var subKegiatan = model.Sub_Kegiatan{
		Tahun:        subKegiatanRequest.Tahun,
		Nama_Program: subKegiatanRequest.Nama_Program,
		Kode:         subKegiatanRequest.Kode,
		Kegiatan:     subKegiatanRequest.Kegiatan,
	}

	newSubKegiatan, err := s.subKegiatanRepository.Create(subKegiatan)

	return newSubKegiatan, err
}

func (s *subKegiatanService) Update(id int, subKegiatanRequest request.UpdateSubKegiatanRequest) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	subKegiatan.Tahun = subKegiatanRequest.Tahun
	subKegiatan.Nama_Program = subKegiatanRequest.Nama_Program
	subKegiatan.Kode = subKegiatanRequest.Kode
	subKegiatan.Kegiatan = subKegiatanRequest.Kegiatan

	updatedSubKegiatan, err := s.subKegiatanRepository.Update(subKegiatan)

	return updatedSubKegiatan, err
}

func (s *subKegiatanService) Delete(id int) (model.Sub_Kegiatan, error) {
	var subKegiatan, err = s.subKegiatanRepository.FindById(id)

	deletedSubKegiatan, err := s.subKegiatanRepository.Delete(subKegiatan)

	return deletedSubKegiatan, err
}
