package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type RincianKwitansiPenerbanganService interface {
	FindAll() ([]model.RincianKwitansiPenerbangan, error)
	FindById(id int) (model.RincianKwitansiPenerbangan, error)
	FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansiPenerbangan, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(rincianKwitansiPenerbanganRequest []request.CreateRincianKwitansiPenerbanganRequest) ([]model.RincianKwitansiPenerbangan, error)
	Update(id int, rincianKwitansiPenerbanganRequest request.UpdateRincianKwitansiPenerbanganRequest) (model.RincianKwitansiPenerbangan, error)
	Delete(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error)
}

type rincianKwitansiPenerbanganService struct {
	rincianKwitansiPenerbanganRepository repository.RincianKwitansiPenerbanganRepository
}

func NewRincianKwitansiPenerbanganService(rincianKwitansiPenerbanganRepository repository.RincianKwitansiPenerbanganRepository) *rincianKwitansiPenerbanganService {
	return &rincianKwitansiPenerbanganService{rincianKwitansiPenerbanganRepository}
}

func (s *rincianKwitansiPenerbanganService) FindAll() ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans, err = s.rincianKwitansiPenerbanganRepository.FindAll()

	return rincianKwitansiPenerbangans, err
}

func (s *rincianKwitansiPenerbanganService) FindById(id int) (model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangan, err = s.rincianKwitansiPenerbanganRepository.FindById(id)

	return rincianKwitansiPenerbangan, err
}

func (s *rincianKwitansiPenerbanganService) FindByKwitansiId(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans, err = s.rincianKwitansiPenerbanganRepository.FindByKwitansiId(kwitansiId)

	return rincianKwitansiPenerbangans, err
}

func (s *rincianKwitansiPenerbanganService) FindBySearch(whereClause map[string]interface{}) ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans, err = s.rincianKwitansiPenerbanganRepository.FindBySearch(whereClause)

	return rincianKwitansiPenerbangans, err
}

func (s *rincianKwitansiPenerbanganService) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count, err = s.rincianKwitansiPenerbanganRepository.CountDataBySearch(whereClause)

	return count, err
}

func (s *rincianKwitansiPenerbanganService) Create(rincianKwitansiPenerbanganRequest []request.CreateRincianKwitansiPenerbanganRequest) ([]model.RincianKwitansiPenerbangan, error) {
	var dataRincianPenerbangans []model.RincianKwitansiPenerbangan

	for _, dataRincianPenerbangan := range rincianKwitansiPenerbanganRequest {
		var data = model.RincianKwitansiPenerbangan{
			Id:           dataRincianPenerbangan.Id,
			KwitansiId:   dataRincianPenerbangan.KwitansiId,
			NamaMaskapai: dataRincianPenerbangan.NamaMaskapai,
			NomorTiket:   dataRincianPenerbangan.NomorTiket,
			KodeBooking:  dataRincianPenerbangan.KodeBooking,
		}

		dataRincianPenerbangans = append(dataRincianPenerbangans, data)
	}

	newRincianKwitansiPenerbangan, err := s.rincianKwitansiPenerbanganRepository.Create(dataRincianPenerbangans)

	return newRincianKwitansiPenerbangan, err
}

func (s *rincianKwitansiPenerbanganService) Update(id int, rincianKwitansiPenerbanganRequest request.UpdateRincianKwitansiPenerbanganRequest) (model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangan, err = s.rincianKwitansiPenerbanganRepository.FindById(id)

	rincianKwitansiPenerbangan.Id = rincianKwitansiPenerbanganRequest.Id
	rincianKwitansiPenerbangan.KwitansiId = rincianKwitansiPenerbanganRequest.KwitansiId
	rincianKwitansiPenerbangan.NamaMaskapai = rincianKwitansiPenerbanganRequest.NamaMaskapai
	rincianKwitansiPenerbangan.NomorTiket = rincianKwitansiPenerbanganRequest.NomorTiket
	rincianKwitansiPenerbangan.KodeBooking = rincianKwitansiPenerbanganRequest.KodeBooking

	updatedRincianKwitansiPenerbangan, err := s.rincianKwitansiPenerbanganRepository.Update(rincianKwitansiPenerbangan)

	return updatedRincianKwitansiPenerbangan, err
}

func (s *rincianKwitansiPenerbanganService) Delete(kwitansiId int) ([]model.RincianKwitansiPenerbangan, error) {
	var rincianKwitansiPenerbangans, err = s.rincianKwitansiPenerbanganRepository.FindByKwitansiId(kwitansiId)

	deletedRincianKwitansiPenerbangan, err := s.rincianKwitansiPenerbanganRepository.Delete(rincianKwitansiPenerbangans)

	return deletedRincianKwitansiPenerbangan, err
}
