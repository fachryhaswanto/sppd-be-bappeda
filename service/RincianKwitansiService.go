package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type RincianKwitansiService interface {
	FindAll() ([]model.RincianKwitansi, error)
	FindById(id int) (model.RincianKwitansi, error)
	FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansi, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(rincianKwitansiRequest []request.CreateRincianKwitansiRequest) ([]model.RincianKwitansi, error)
	Update(id int, rincianKwitansiRequest request.UpdateRincianKwitansiRequest) (model.RincianKwitansi, error)
	Delete(kwitansiId int) ([]model.RincianKwitansi, error)
}

type rincianKwitansiService struct {
	rincianKwitansiRepository repository.RincianKwitansiRepository
}

func NewRincianKwitansiService(rincianKwitansiRepository repository.RincianKwitansiRepository) *rincianKwitansiService {
	return &rincianKwitansiService{rincianKwitansiRepository}
}

func (s *rincianKwitansiService) FindAll() ([]model.RincianKwitansi, error) {
	var rincianKwitansis, err = s.rincianKwitansiRepository.FindAll()

	return rincianKwitansis, err
}

func (s *rincianKwitansiService) FindById(id int) (model.RincianKwitansi, error) {
	var rincianKwitansi, err = s.rincianKwitansiRepository.FindById(id)

	return rincianKwitansi, err
}

func (s *rincianKwitansiService) FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansi, error) {
	var rincianKwitansis, err = s.rincianKwitansiRepository.FindByKwitansiId(kwitansiId)

	return rincianKwitansis, err
}

func (s *rincianKwitansiService) FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansi, error) {
	var rincianKwitansis, err = s.rincianKwitansiRepository.FindBySearch(whereClause)

	return rincianKwitansis, err
}

func (s *rincianKwitansiService) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count, err = s.rincianKwitansiRepository.CountDataBySearch(whereClause)

	return count, err
}

func (s *rincianKwitansiService) Create(rincianKwitansiRequest []request.CreateRincianKwitansiRequest) ([]model.RincianKwitansi, error) {
	var dataRincians []model.RincianKwitansi

	for _, dataRincian := range rincianKwitansiRequest {
		var data = model.RincianKwitansi{
			Id:          dataRincian.Id,
			KwitansiId:  dataRincian.KwitansiId,
			NamaRincian: dataRincian.NamaRincian,
			JumlahBayar: dataRincian.JumlahBayar,
			Banyaknya:   dataRincian.Banyaknya,
			HasilBayar:  dataRincian.HasilBayar,
		}

		dataRincians = append(dataRincians, data)
	}

	newRincianKwitansi, err := s.rincianKwitansiRepository.Create(dataRincians)

	return newRincianKwitansi, err
}

func (s *rincianKwitansiService) Update(id int, rincianKwitansiRequest request.UpdateRincianKwitansiRequest) (model.RincianKwitansi, error) {
	var rincianKwitansi, err = s.rincianKwitansiRepository.FindById(id)

	rincianKwitansi.Id = rincianKwitansiRequest.Id
	rincianKwitansi.KwitansiId = rincianKwitansiRequest.KwitansiId
	rincianKwitansi.NamaRincian = rincianKwitansiRequest.NamaRincian
	rincianKwitansi.JumlahBayar = rincianKwitansiRequest.JumlahBayar
	rincianKwitansi.Banyaknya = rincianKwitansiRequest.Banyaknya
	rincianKwitansi.HasilBayar = rincianKwitansiRequest.HasilBayar

	updatedRincianKwitansi, err := s.rincianKwitansiRepository.Update(rincianKwitansi)

	return updatedRincianKwitansi, err
}

func (s *rincianKwitansiService) Delete(kwitansiId int) ([]model.RincianKwitansi, error) {
	var rincianKwitansis, err = s.rincianKwitansiRepository.FindByKwitansiId(kwitansiId)

	deletedRincianKwitansi, err := s.rincianKwitansiRepository.Delete(rincianKwitansis)

	return deletedRincianKwitansi, err
}
