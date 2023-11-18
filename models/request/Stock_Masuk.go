package request

type Input_Stock_Masuk_Request struct {
	Co                    int    `json:"co"`
	Kode_stock_masuk      string `json:"kode_stock_masuk"`
	Tanggal_masuk         string `json:"Tanggal_masuk"`
	Kode_nota             string `json:"kode_nota"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Kode_supplier         string `json:"kode_supplier"`
	Kode_gudang           string `json:"kode_gudang"`
}

type Input_Barang_Stock_Masuk_Request struct {
	Kode_stock         string `json:"kode_stock"`
	Tanggal_kadalurasa string `json:"tanggal_kadalurasa"`
	Jumlah_barang      string `json:"jumlah_barang"`
	Harga_pokok        string `json:"harga_pokok"`
}

type Input_Barang_Stock_Masuk_V2_Request struct {
	Co                       int     `json:"co"`
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Kode_stock_keluar_masuk  string  `json:"kode_stock_keluar_masuk"`
	Kode_stock               string  `json:"kode_stock"`
	Tanggal_kadaluarsa       string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang            float64 `json:"jumlah_barang"`
	Harga                    int64   `json:"harga"`
	Total_harga              int64   `json:"total_harga"`
	Status                   int     `json:" status"`
}

type Read_Stock_Masuk_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Stock_Masuk_Filter_Request struct {
	Tanggal_1     string `json:"tanggal_1"`
	Tanggal_2     string `json:"tanggal_2"`
	Kode_supplier string `json:"kode_supplier"`
}
