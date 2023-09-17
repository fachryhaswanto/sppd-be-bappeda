package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type KwitansiService interface {
	FindAll() ([]model.Kwitansi, error)
	FindById(id int) (model.Kwitansi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Kwitansi, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	SumTotalBayar(whereClause map[string]interface{}) int64
	Create(kwitansiRequest request.CreateKwitansiRequest) (model.Kwitansi, error)
	Update(id int, kwitansiRequest request.UpdateKwitansiRequest) (model.Kwitansi, error)
	UpdateTotalBayar(id int, kwitansiRequest request.UpdateTotalBayarRequest) (model.Kwitansi, error)
	Delete(id int) (model.Kwitansi, error)
}

type kwitansiService struct {
	kwitansiRepository repository.KwitansiRepository
}

func NewKwitansiService(kwitansiRepository repository.KwitansiRepository) *kwitansiService {
	return &kwitansiService{kwitansiRepository}
}

func (s *kwitansiService) FindAll() ([]model.Kwitansi, error) {
	var kwitansis, err = s.kwitansiRepository.FindAll()

	return kwitansis, err
}

func (s *kwitansiService) FindById(id int) (model.Kwitansi, error) {
	var kwitansi, err = s.kwitansiRepository.FindById(id)

	return kwitansi, err
}

func (s *kwitansiService) FindBySearch(whereClause map[string]interface{}) ([]model.Kwitansi, error) {
	var kwitansis, err = s.kwitansiRepository.FindBySearch(whereClause)

	return kwitansis, err
}

func (s *kwitansiService) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count, err = s.kwitansiRepository.CountDataBySearch(whereClause)

	return count, err
}

func (s *kwitansiService) SumTotalBayar(whereClause map[string]interface{}) int64 {
	var realisasi = s.kwitansiRepository.SumTotalBayar(whereClause)

	return realisasi
}

func (s *kwitansiService) Create(kwitansiRequest request.CreateKwitansiRequest) (model.Kwitansi, error) {
	var tahun = kwitansiRequest.TanggalBayar[:4]

	var kwitansi = model.Kwitansi{
		SppdId:        kwitansiRequest.SppdId,
		NomorKwitansi: kwitansiRequest.NomorKwitansi,
		TanggalBayar:  kwitansiRequest.TanggalBayar,
		Keperluan:     kwitansiRequest.Keperluan,
		TotalBayar:    0,
		Tahun:         tahun,
		UserId:        kwitansiRequest.UserId,
	}

	var newKwitansi, err = s.kwitansiRepository.Create(kwitansi)

	return newKwitansi, err
}

func (s *kwitansiService) Update(id int, kwitansiRequest request.UpdateKwitansiRequest) (model.Kwitansi, error) {
	var tahun = kwitansiRequest.TanggalBayar[:4]
	var kwitansi, err = s.kwitansiRepository.FindById(id)

	kwitansi.SppdId = kwitansiRequest.SppdId
	kwitansi.NomorKwitansi = kwitansiRequest.NomorKwitansi
	kwitansi.TanggalBayar = kwitansiRequest.TanggalBayar
	kwitansi.Keperluan = kwitansiRequest.Keperluan
	kwitansi.TotalBayar = kwitansiRequest.TotalBayar
	kwitansi.Tahun = tahun
	kwitansi.UserId = kwitansiRequest.UserId

	updatedKwitansi, err := s.kwitansiRepository.Update(kwitansi)

	return updatedKwitansi, err
}

func (s *kwitansiService) UpdateTotalBayar(id int, kwitansiRequest request.UpdateTotalBayarRequest) (model.Kwitansi, error) {
	var kwitansi, err = s.kwitansiRepository.FindById(id)

	kwitansi.TotalBayar = kwitansiRequest.TotalBayar

	updatedKwitansi, err := s.kwitansiRepository.UpdateTotalBayar(kwitansi)

	return updatedKwitansi, err

}

func (s *kwitansiService) Delete(id int) (model.Kwitansi, error) {
	var kwitansi, err = s.kwitansiRepository.FindById(id)

	deletedSpt, err := s.kwitansiRepository.Delete(kwitansi)

	return deletedSpt, err
}
