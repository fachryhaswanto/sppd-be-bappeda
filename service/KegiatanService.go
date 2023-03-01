package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type KegiatanService interface {
	FindAll() ([]model.Kegiatan, error)
	FindById(id int) (model.Kegiatan, error)
	Create(kegiatanRequest request.CreateKegiatanRequest) (model.Kegiatan, error)
	Update(id int, kegiatanRequest request.UpdateKegiatanRequest) (model.Kegiatan, error)
	Delete(id int) (model.Kegiatan, error)
}

type kegiatanService struct {
	kegiatanRepository repository.KegiatanRepository
}

func NewKegiatanService(kegiatanRepository repository.KegiatanRepository) *kegiatanService {
	return &kegiatanService{kegiatanRepository}
}

func (s *kegiatanService) FindAll() ([]model.Kegiatan, error) {
	var kegiatans, err = s.kegiatanRepository.FindAll()

	return kegiatans, err
}

func (s *kegiatanService) FindById(id int) (model.Kegiatan, error) {
	var kegiatan, err = s.kegiatanRepository.FindById(id)

	return kegiatan, err
}

func (s *kegiatanService) Create(kegiatanRequest request.CreateKegiatanRequest) (model.Kegiatan, error) {
	var kegiatan = model.Kegiatan{
		ProgramId:    kegiatanRequest.ProgramId,
		KodeKegiatan: kegiatanRequest.KodeKegiatan,
		NamaKegiatan: kegiatanRequest.NamaKegiatan,
	}

	var newKegiatan, err = s.kegiatanRepository.Create(kegiatan)

	return newKegiatan, err
}

func (s *kegiatanService) Update(id int, kegiatanRequest request.UpdateKegiatanRequest) (model.Kegiatan, error) {
	var kegiatan, err = s.kegiatanRepository.FindById(id)

	kegiatan.ProgramId = kegiatanRequest.ProgramId
	kegiatan.KodeKegiatan = kegiatanRequest.KodeKegiatan
	kegiatan.NamaKegiatan = kegiatanRequest.NamaKegiatan

	updatedKegiatan, err := s.kegiatanRepository.Update(kegiatan)

	return updatedKegiatan, err
}

func (s *kegiatanService) Delete(id int) (model.Kegiatan, error) {
	var kegiatan, err = s.kegiatanRepository.FindById(id)

	deletedKegiatan, err := s.kegiatanRepository.Delete(kegiatan)

	return deletedKegiatan, err
}
