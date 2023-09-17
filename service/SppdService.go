package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type SppdService interface {
	FindAll() ([]model.Sppd, error)
	FindById(id int) (model.Sppd, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Sppd, error)
	CountDataBySptId(sptId int) (int64, error)
	CountDataBySearch(whereClause map[string]interface{}) (int64, error)
	Create(sppdRequest request.CreateSppdRequest) (model.Sppd, error)
	Update(id int, sppd request.UpdateSppdRequest) (model.Sppd, error)
	UpdateBySptId(sptId int, sppd request.UpdateSppdRequest) (model.Sppd, error)
	UpdateStatusKwitansi(id, value int) error
	Delete(id int) (model.Sppd, error)
	DeleteBySptId(sptId int) error
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

func (s *sppdService) FindById(id int) (model.Sppd, error) {
	var sppd, err = s.sppdRepository.FindById(id)

	return sppd, err
}

func (s *sppdService) FindBySearch(whereClause map[string]interface{}) ([]model.Sppd, error) {
	var sppds, err = s.sppdRepository.FindBySearch(whereClause)

	return sppds, err
}

func (s *sppdService) CountDataBySptId(sptId int) (int64, error) {
	var count, err = s.sppdRepository.CountDataBySptId(sptId)

	return count, err
}

func (s *sppdService) CountDataBySearch(whereClause map[string]interface{}) (int64, error) {
	var count, err = s.sppdRepository.CountDataBySearch(whereClause)

	return count, err
}

func (s *sppdService) Create(sppdRequest request.CreateSppdRequest) (model.Sppd, error) {
	var tahun = sppdRequest.Tanggal_Sppd[:4]

	var sppd = model.Sppd{
		Template_Sppd:    sppdRequest.Template_Sppd,
		Jenis:            sppdRequest.Jenis,
		Nomor_Sppd:       sppdRequest.Nomor_Sppd,
		PegawaiId:        sppdRequest.PegawaiId,
		Tanggal_Sppd:     sppdRequest.Tanggal_Sppd,
		Tempat_Berangkat: sppdRequest.Tempat_Berangkat,
		Tempat_Tujuan:    sppdRequest.Tempat_Tujuan,
		Tahun:            tahun,
		Alat_Angkutan:    sppdRequest.Alat_Angkutan,
		Instansi:         sppdRequest.Instansi,
		PejabatId:        sppdRequest.PejabatId,
		StatusKwitansi:   0,
		SptId:            sppdRequest.SptId,
		UserId:           sppdRequest.UserId,
	}

	newSppd, err := s.sppdRepository.Create(sppd)

	return newSppd, err
}

func (s *sppdService) Update(id int, sppdRequest request.UpdateSppdRequest) (model.Sppd, error) {
	var tahun = sppdRequest.Tanggal_Sppd[:4]
	var sppd, err = s.sppdRepository.FindById(id)

	sppd.Template_Sppd = sppdRequest.Template_Sppd
	sppd.Jenis = sppdRequest.Jenis
	sppd.Nomor_Sppd = sppdRequest.Nomor_Sppd
	sppd.PegawaiId = sppdRequest.PegawaiId
	sppd.Tanggal_Sppd = sppdRequest.Tanggal_Sppd
	sppd.Tempat_Berangkat = sppdRequest.Tempat_Berangkat
	sppd.Tempat_Tujuan = sppdRequest.Tempat_Tujuan
	sppd.Tahun = tahun
	sppd.Alat_Angkutan = sppdRequest.Alat_Angkutan
	sppd.Instansi = sppdRequest.Instansi
	sppd.PejabatId = sppdRequest.PejabatId
	sppd.StatusKwitansi = sppdRequest.StatusKwitansi
	sppd.SptId = sppdRequest.SptId
	sppd.UserId = sppdRequest.UserId

	updatedSppd, err := s.sppdRepository.Update(sppd)

	return updatedSppd, err
}

func (s *sppdService) UpdateBySptId(sptId int, sppdRequest request.UpdateSppdRequest) (model.Sppd, error) {
	var sppd = model.Sppd{
		Template_Sppd:    sppdRequest.Template_Sppd,
		Jenis:            sppdRequest.Jenis,
		Nomor_Sppd:       sppdRequest.Nomor_Sppd,
		PegawaiId:        sppdRequest.PegawaiId,
		Tanggal_Sppd:     sppdRequest.Tanggal_Sppd,
		Tempat_Berangkat: sppdRequest.Tempat_Berangkat,
		Tempat_Tujuan:    sppdRequest.Tempat_Tujuan,
		Alat_Angkutan:    sppdRequest.Alat_Angkutan,
		Instansi:         sppdRequest.Instansi,
		PejabatId:        sppdRequest.PejabatId,
		StatusKwitansi:   sppdRequest.StatusKwitansi,
		SptId:            sppdRequest.SptId,
		UserId:           sppdRequest.UserId,
	}

	updatedSppd, err := s.sppdRepository.UpdateBySptId(sptId, sppd)

	return updatedSppd, err
}

func (s *sppdService) UpdateStatusKwitansi(id, value int) error {
	var err = s.sppdRepository.UpdateStatusKwitansi(id, value)

	return err
}

func (s *sppdService) Delete(id int) (model.Sppd, error) {
	var sppd, err = s.sppdRepository.FindById(id)

	deletedSppd, err := s.sppdRepository.Delete(sppd)

	return deletedSppd, err
}

func (s *sppdService) DeleteBySptId(sptId int) error {
	var err = s.sppdRepository.DeleteBySptId(sptId)

	return err
}
