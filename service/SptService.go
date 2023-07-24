package service

import (
	"sppd/model"
	"sppd/repository"
	"sppd/request"
)

type SptService interface {
	FindAll() ([]model.Spt, error)
	FindById(id int) (model.Spt, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Spt, error)
	FindAllDitugaskan() ([]model.Spt, error)
	Create(spt request.CreateSptRequest) (model.Spt, error)
	Update(id int, spt request.UpdateSptRequest) (model.Spt, error)
	UpdateStatusSppd(id int, value int) error
	Delete(id int) (model.Spt, error)
}

type sptService struct {
	sptRepository repository.SptRepository
}

func NewSptService(repository repository.SptRepository) *sptService {
	return &sptService{repository}
}

func (s *sptService) FindAll() ([]model.Spt, error) {
	var spts, err = s.sptRepository.FindAll()

	return spts, err
}

func (s *sptService) FindById(id int) (model.Spt, error) {
	var spt, err = s.sptRepository.FindById(id)

	return spt, err
}

func (s *sptService) FindBySearch(whereClause map[string]interface{}) ([]model.Spt, error) {
	var spts, err = s.sptRepository.FindBySearch(whereClause)

	return spts, err
}

func (s *sptService) FindAllDitugaskan() ([]model.Spt, error) {
	var ditugaskan, err = s.sptRepository.FindAllDitugaskan()

	return ditugaskan, err
}

func (s *sptService) Create(sptRequest request.CreateSptRequest) (model.Spt, error) {
	var spt = model.Spt{
		Jenis:             sptRequest.Jenis,
		Template:          sptRequest.Template,
		SubKegiatanId:     sptRequest.SubKegiatanId,
		Nomor_Spt:         sptRequest.Nomor_Spt,
		Tanggal_Spt:       sptRequest.Tanggal_Spt,
		RekeningId:        sptRequest.RekeningId,
		Keperluan:         sptRequest.Keperluan,
		Tanggal_Berangkat: sptRequest.Tanggal_Berangkat,
		Tanggal_Kembali:   sptRequest.Tanggal_Kembali,
		Lama_Perjalanan:   sptRequest.Lama_Perjalanan,
		Pejabat_Pemberi:   sptRequest.Pejabat_Pemberi,
		Status:            "Belum Kembali",
		StatusSppd:        0,
		File_Surat_Tugas:  sptRequest.File_Surat_Tugas,
		UserId:            sptRequest.UserId,
	}

	var newSpt, err = s.sptRepository.Create(spt)

	return newSpt, err
}

func (s *sptService) Update(id int, sptRequest request.UpdateSptRequest) (model.Spt, error) {
	var spt, err = s.sptRepository.FindById(id)

	spt.Jenis = sptRequest.Jenis
	spt.Template = sptRequest.Template
	spt.SubKegiatanId = sptRequest.SubKegiatanId
	spt.Nomor_Spt = sptRequest.Nomor_Spt
	spt.Tanggal_Spt = sptRequest.Tanggal_Spt
	spt.RekeningId = sptRequest.RekeningId
	spt.Keperluan = sptRequest.Keperluan
	spt.Tanggal_Berangkat = sptRequest.Tanggal_Berangkat
	spt.Tanggal_Kembali = sptRequest.Tanggal_Kembali
	spt.Lama_Perjalanan = sptRequest.Lama_Perjalanan
	spt.Pejabat_Pemberi = sptRequest.Pejabat_Pemberi
	spt.Status = sptRequest.Status
	spt.StatusSppd = sptRequest.StatusSppd
	spt.File_Surat_Tugas = sptRequest.File_Surat_Tugas
	spt.UserId = sptRequest.UserId

	updatedSpt, err := s.sptRepository.Update(spt)

	return updatedSpt, err
}

func (s *sptService) UpdateStatusSppd(id int, value int) error {
	var err = s.sptRepository.UpdateStatusSppd(id, value)

	return err
}

func (s *sptService) Delete(id int) (model.Spt, error) {
	var spt, err = s.sptRepository.FindById(id)

	deletedSpt, err := s.sptRepository.Delete(spt)

	return deletedSpt, err
}
