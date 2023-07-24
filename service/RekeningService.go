package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type RekeningService interface {
	FindAll() ([]model.Rekening, error)
	FindById(id int) (model.Rekening, error)
	Create(rekeningRequest request.CreateRekeningRequest) (model.Rekening, error)
	Update(id int, rekeningRequest request.UpdateRekeningRequest) (model.Rekening, error)
	Delete(id int) (model.Rekening, error)
}

type rekeningService struct {
	rekeningRepository repository.RekeningRepository
}

func NewRekeningService(rekeningRepository repository.RekeningRepository) *rekeningService {
	return &rekeningService{rekeningRepository}
}

func (s *rekeningService) FindAll() ([]model.Rekening, error) {
	var rekenings, err = s.rekeningRepository.FindAll()

	return rekenings, err
}

func (s *rekeningService) FindById(id int) (model.Rekening, error) {
	var rekening, err = s.rekeningRepository.FindById(id)

	return rekening, err
}

func (s *rekeningService) Create(rekeningRequest request.CreateRekeningRequest) (model.Rekening, error) {
	var rekening = model.Rekening{
		KodeRekening: rekeningRequest.KodeRekening,
		NamaRekening: rekeningRequest.NamaRekening,
	}

	newRekening, err := s.rekeningRepository.Create(rekening)

	return newRekening, err
}

func (s *rekeningService) Update(id int, rekeningRequest request.UpdateRekeningRequest) (model.Rekening, error) {
	var rekening, err = s.rekeningRepository.FindById(id)

	rekening.KodeRekening = rekeningRequest.KodeRekening
	rekening.NamaRekening = rekeningRequest.NamaRekening

	updatedRekening, err := s.rekeningRepository.Update(rekening)

	return updatedRekening, err
}

func (s *rekeningService) Delete(id int) (model.Rekening, error) {
	var rekening, err = s.rekeningRepository.FindById(id)

	deletedRekening, err := s.rekeningRepository.Delete(rekening)

	return deletedRekening, err
}
