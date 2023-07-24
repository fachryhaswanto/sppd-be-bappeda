package service

import (
	"errors"
	"sppd/helper"
	"sppd/model"
	"sppd/repository"
	"sppd/request"
	"strconv"
	"strings"
	"time"
)

type DataDitugaskanService interface {
	FindAll() ([]model.DataDitugaskan, error)
	FindById(id int) (model.DataDitugaskan, error)
	FindBySptId(sptId int) ([]model.DataDitugaskan, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.DataDitugaskan, error)
	CountDataBySptId(sptId int) (int64, error)
	Create(dataDitugaskanRequest []request.CreateDataDitugaskanRequest) ([]model.DataDitugaskan, error)
	UpdateStatus(sptId int, pegawaiId int, value int) error
	Delete(sptId int) ([]model.DataDitugaskan, error)
}

type dataDitugaskanService struct {
	dataDitugaskanRepository repository.DataDitugaskanRepository
	sptRepository            repository.SptRepository
}

func NewDataDitugaskanService(dataDitugaskanRepository repository.DataDitugaskanRepository, sptRepository repository.SptRepository) *dataDitugaskanService {
	return &dataDitugaskanService{dataDitugaskanRepository, sptRepository}
}

func (s *dataDitugaskanService) FindAll() ([]model.DataDitugaskan, error) {
	var dataDitugaskans, err = s.dataDitugaskanRepository.FindAll()

	return dataDitugaskans, err
}

func (s *dataDitugaskanService) FindById(id int) (model.DataDitugaskan, error) {
	var dataDitugaskan, err = s.dataDitugaskanRepository.FindById(id)

	return dataDitugaskan, err
}

func (s *dataDitugaskanService) FindBySptId(sptId int) ([]model.DataDitugaskan, error) {
	var dataDitugaskans, err = s.dataDitugaskanRepository.FindBySptId(sptId)

	return dataDitugaskans, err
}

func (s *dataDitugaskanService) FindBySearch(whereClause map[string]interface{}) ([]model.DataDitugaskan, error) {
	var dataDitugaskans, err = s.dataDitugaskanRepository.FindBySearch(whereClause)

	return dataDitugaskans, err
}

func (s *dataDitugaskanService) CountDataBySptId(sptId int) (int64, error) {
	var count, err = s.dataDitugaskanRepository.CountDataBySptId(sptId)

	return count, err
}

func (s *dataDitugaskanService) Create(dataDitugaskanRequests []request.CreateDataDitugaskanRequest) ([]model.DataDitugaskan, error) {
	var dataDitugaskans []model.DataDitugaskan
	var dataDitugaskansFilteredByDate []model.DataDitugaskan
	var errMessage []string

	var jumlahDataSpt, err = s.sptRepository.CountDataSpt()

	if jumlahDataSpt == 1 {
		// fmt.Println("data spt hanya 1 tidak perlu dilakukan pengecekan")
	} else {
		var dataSpt, _ = s.sptRepository.FindById(dataDitugaskanRequests[0].SptId)

		var tanggalBerangkat = dataSpt.Tanggal_Berangkat
		var tanggalKembali = dataSpt.Tanggal_Kembali

		var tanggalBerangkatSplit = strings.Split(tanggalBerangkat, "/")
		var tanggalKembaliSplit = strings.Split(tanggalKembali, "/")

		var dataSptFilteredByDate, _ = s.sptRepository.FindByDate(tanggalBerangkatSplit[0], tanggalBerangkatSplit[1], tanggalKembaliSplit[0], tanggalKembaliSplit[1], dataDitugaskanRequests[0].SptId)

		for _, data := range dataSptFilteredByDate {
			if data.Id != dataDitugaskanRequests[0].SptId {
				var dataDitugaskan, _ = s.dataDitugaskanRepository.FindBySptId(data.Id)
				dataDitugaskansFilteredByDate = append(dataDitugaskansFilteredByDate, dataDitugaskan...)
			}
		}

		for _, dataRequest := range dataDitugaskanRequests {
			for _, dataFiltered := range dataDitugaskansFilteredByDate {
				if dataRequest.PegawaiId == dataFiltered.PegawaiId {

					var dataSptByPegawaiId, _ = s.sptRepository.FindById(dataFiltered.SptId)

					var tanggalBerangkatPembanding = strings.Split(dataSptByPegawaiId.Tanggal_Berangkat, "/")
					var tanggalKembaliPembanding = strings.Split(dataSptByPegawaiId.Tanggal_Kembali, "/")

					var tahunBerangkat, _ = strconv.Atoi(tanggalBerangkatPembanding[0])
					var bulanBerangkat = helper.ConvertToTimeMonth(tanggalBerangkatPembanding[1])
					var tanggalBerangkat, _ = strconv.Atoi(tanggalBerangkatPembanding[2])

					var tahunKembali, _ = strconv.Atoi(tanggalKembaliPembanding[0])
					var bulanKembali = helper.ConvertToTimeMonth(tanggalKembaliPembanding[1])
					var tanggalKembali, _ = strconv.Atoi(tanggalKembaliPembanding[2])

					var tahunTargetBerangkat, _ = strconv.Atoi(tanggalBerangkatSplit[0])
					var bulanTargetBerangkat = helper.ConvertToTimeMonth(tanggalBerangkatSplit[1])
					var tanggalTargetBerangkat, _ = strconv.Atoi(tanggalBerangkatSplit[2])

					var tahunTargetKembali, _ = strconv.Atoi(tanggalKembaliSplit[0])
					var bulanTargetKembali = helper.ConvertToTimeMonth(tanggalKembaliSplit[1])
					var tanggalTargetKembali, _ = strconv.Atoi(tanggalKembaliSplit[2])

					var startDate = time.Date(tahunBerangkat, bulanBerangkat, tanggalBerangkat, 0, 0, 0, 0, time.UTC)
					var endDate = time.Date(tahunKembali, bulanKembali, tanggalKembali, 0, 0, 0, 0, time.UTC)

					var targetDateBerangkat = time.Date(tahunTargetBerangkat, bulanTargetBerangkat, tanggalTargetBerangkat, 0, 0, 0, 0, time.UTC)
					var targetDateKembali = time.Date(tahunTargetKembali, bulanTargetKembali, tanggalTargetKembali, 0, 0, 0, 0, time.UTC)

					if dataSpt.Jenis == "Baru" {
						if (targetDateBerangkat.Equal(startDate) || targetDateBerangkat.Equal(endDate)) || (targetDateBerangkat.After(startDate) && targetDateBerangkat.Before(endDate)) {
							var msg = "pegawai dengan nama : " + dataFiltered.Pegawai.Nama + " memiliki kesamaan tanggal pada tanggal berangkat atau tanggal kembali ataupun tanggal berangkat berada dalam jangkauan tanggal berangkat dan tanggal kembali pada spt dengan nomor : " + dataSptByPegawaiId.Nomor_Spt
							errMessage = append(errMessage, msg)
						}
					} else {
						if targetDateBerangkat.Equal(startDate) || (targetDateBerangkat.After(startDate) && targetDateBerangkat.Before(endDate)) {
							var msg = "pegawai dengan nama : " + dataFiltered.Pegawai.Nama + " memiliki kesamaan tanggal pada tanggal berangkat atau tanggal berangkat berada dalam jangkauan tanggal berangkat dan tanggal kembali pada spt dengan nomor : " + dataSptByPegawaiId.Nomor_Spt
							errMessage = append(errMessage, msg)
						}
					}

					if targetDateKembali.Equal(endDate) || (targetDateKembali.After(startDate) && targetDateKembali.Before(endDate)) {
						var msg = "pegawai dengan nama : " + dataFiltered.Pegawai.Nama + " memiliki kesamaan tanggal pada tanggal kembali atau tanggal kembali berada dalam jangkauan tanggal berangkat dan tanggal kembali pada spt dengan nomor :" + dataSptByPegawaiId.Nomor_Spt
						errMessage = append(errMessage, msg)
					}
				}
			}
		}
	}

	for _, dataDitugaskan := range dataDitugaskanRequests {
		var data = model.DataDitugaskan{
			SptId:      dataDitugaskan.SptId,
			PegawaiId:  dataDitugaskan.PegawaiId,
			StatusSppd: 0,
		}
		dataDitugaskans = append(dataDitugaskans, data)
	}

	var concatenatedErrString = strings.Join(errMessage, ", ")

	if len(errMessage) != 0 {
		return dataDitugaskans, errors.New(concatenatedErrString)
	}

	newDataDitugaskan, err := s.dataDitugaskanRepository.Create(dataDitugaskans)

	return newDataDitugaskan, err
}

func (s *dataDitugaskanService) UpdateStatus(sptId int, pegawaiId int, value int) error {
	var err = s.dataDitugaskanRepository.UpdateStatus(sptId, pegawaiId, value)

	return err
}

func (s *dataDitugaskanService) Delete(sptId int) ([]model.DataDitugaskan, error) {
	var dataDitugaskans, err = s.dataDitugaskanRepository.FindBySptId(sptId)

	deletedDataDitugaskan, err := s.dataDitugaskanRepository.Delete(dataDitugaskans)

	return deletedDataDitugaskan, err
}
