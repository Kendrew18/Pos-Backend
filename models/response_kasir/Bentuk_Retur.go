package response_kasir

type Read_Bentuk_Retur_Response struct {
	Co                int    `json:"co"`
	Kode_bentuk_retur string `json:"kode_bentuk_retur"`
	Nama_bentuk_retur string `json:"nama_bentuk_retur"`
	Nama_satuan       string `json:"nama_satuan"`
}
