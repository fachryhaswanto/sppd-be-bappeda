package request

type CreateRekeningRequest struct {
	KodeRekening string `json:"kodeRekening" binding:"required"`
	NamaRekening string `json:"namaRekening" binding:"required"`
}

type UpdateRekeningRequest struct {
	KodeRekening string `json:"kodeRekening"`
	NamaRekening string `json:"namaRekening"`
}
