package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type SppdService interface {
	FindAll() ([]model.Sppd, error)
	FindJoin() ([]model.JoinSppdSpt, error)
	FindById(id int) (model.Sppd, error)
	Create(sppd request.CreateSppdRequest) (model.Sppd, error)
	Update(id int, sppd request.UpdateSppdRequest) (model.Sppd, error)
	Delete(id int) (model.Sppd, error)
}

type sppdService struct {
	sppdRepository repository.SppdRepository
}

func NewSppdService(repository repository.SppdRepository) *sppdService {
	return &sppdService{repository}
}

func (s *sppdService) FindAll() ([]model.Sppd, error) {
	var sppds, err = s.sppdRepository.FindAll()

	return sppds, err
}

func (s *sppdService) FindJoin() ([]model.JoinSppdSpt, error) {
	var joins, err = s.sppdRepository.FindJoin()

	return joins, err
}

func (s *sppdService) FindById(id int) (model.Sppd, error) {
	var sppd, err = s.sppdRepository.FindById(id)

	return sppd, err
}

func (s *sppdService) Create(sppdRequest request.CreateSppdRequest) (model.Sppd, error) {
	var sppd = model.Sppd{
		Template_Sppd: sppdRequest.Template_Sppd,
		Nomor_Sppd:    sppdRequest.Nomor_Sppd,
		Tanggal_Sppd:  sppdRequest.Tanggal_Sppd,
		Tingkat_Biaya: sppdRequest.Tingkat_Biaya,
		Instansi:      sppdRequest.Instansi,
		Tanda_Tangan:  sppdRequest.Tanda_tangan,
	}

	newSppd, err := s.sppdRepository.Create(sppd)

	return newSppd, err
}

func (s *sppdService) Update(id int, sppdRequest request.UpdateSppdRequest) (model.Sppd, error) {
	var sppd, err = s.sppdRepository.FindById(id)

	sppd.Template_Sppd = sppdRequest.Template_Sppd
	sppd.Nomor_Sppd = sppdRequest.Nomor_Sppd
	sppd.Tanggal_Sppd = sppdRequest.Tanggal_Sppd
	sppd.Tingkat_Biaya = sppdRequest.Tingkat_Biaya
	sppd.Instansi = sppdRequest.Instansi
	sppd.Tanda_Tangan = sppdRequest.Tanda_tangan
	sppd.Idspt = sppdRequest.Idspt

	updatedSppd, err := s.sppdRepository.Update(sppd)

	return updatedSppd, err
}

func (s *sppdService) Delete(id int) (model.Sppd, error) {
	var sppd, err = s.sppdRepository.FindById(id)

	deletedSppd, err := s.sppdRepository.Delete(sppd)

	return deletedSppd, err
}
