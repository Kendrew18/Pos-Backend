package response

type Read_Stock_Masuk_Response struct {
	Kode_stock_keluar_masuk string                             `json:"kode_stock_keluar_masuk"`
	Tanggal                 string                             `json:"tanggal"`
	Kode_nota               string                             `json:"kode_nota"`
	Penanggung_jawab        string                             `json:"penanggung_jawab"`
	Nama_supplier           string                             `json:"nama_supplier"`
	Jumlah_total            float64                            `json:"jumlah"`
	Total_harga             int64                              `json:"total_harga"`
	Detail_stock_masuk      []Read_Detail_Stock_Masuk_Response `json:"detail_stock_masuk"`
}

type Read_Detail_Stock_Masuk_Response struct {
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Nama_barang              string  `json:"nama_barang"`
	Tanggal_kadaluarsa       string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang            float64 `json:"jumlah_barang"`
	Harga                    int64   `json:"harga"`
	Status                   int     `json:"status"`
}

type Stock_Masuk_Lama_Response struct {
	Kode_stock_keluar_masuk string  `json:"kode_stock_keluar_masuk"`
	Kode_stock              string  `json:"kode_stock"`
	Jumlah_barang           float64 `json:"jumlah_barang"`
}
