package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type DataPengikutService interface {
	FindAll() ([]model.DataPengikut, error)
	FindById(id int) (model.DataPengikut, error)
	FindBySptId(sptId int) ([]model.DataPengikut, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.DataPengikut, error)
	Create(dataPengikutRequest []request.CreateDataPengikutRequest) ([]model.DataPengikut, error)
	Delete(sptId int) ([]model.DataPengikut, error)
}

type dataPengikutService struct {
	dataPengikutRepository repository.DataPengikutRepository
}

func NewDataPengikutService(dataPengikutRepository repository.DataPengikutRepository) *dataPengikutService {
	return &dataPengikutService{dataPengikutRepository}
}

func (s *dataPengikutService) FindAll() ([]model.DataPengikut, error) {
	var dataPengikuts, err = s.dataPengikutRepository.FindAll()

	return dataPengikuts, err
}

func (s *dataPengikutService) FindById(id int) (model.DataPengikut, error) {
	var dataPengikut, err = s.dataPengikutRepository.FindById(id)

	return dataPengikut, err
}

func (s *dataPengikutService) FindBySptId(sptId int) ([]model.DataPengikut, error) {
	var dataPengikuts, err = s.dataPengikutRepository.FindBySptId(sptId)

	return dataPengikuts, err
}

func (s *dataPengikutService) FindBySearch(whereClause map[string]interface{}) ([]model.DataPengikut, error) {
	var dataPengikuts, err = s.dataPengikutRepository.FindBySearch(whereClause)

	return dataPengikuts, err
}

func (s *dataPengikutService) Create(dataPengikutRequest []request.CreateDataPengikutRequest) ([]model.DataPengikut, error) {
	var dataPengikuts []model.DataPengikut

	for _, dataPengikut := range dataPengikutRequest {
		var data = model.DataPengikut{
			SptId:     dataPengikut.SptId,
			PegawaiId: dataPengikut.PegawaiId,
		}

		dataPengikuts = append(dataPengikuts, data)
	}

	newDataPengikuts, err := s.dataPengikutRepository.Create(dataPengikuts)

	return newDataPengikuts, err
}

func (s *dataPengikutService) Delete(sptId int) ([]model.DataPengikut, error) {
	var dataPengikuts, err = s.dataPengikutRepository.FindBySptId(sptId)

	deletedDataPengikuts, err := s.dataPengikutRepository.Delete(dataPengikuts)

	return deletedDataPengikuts, err
}
