package request

type Input_Pre_Order_Request struct {
	Co                    int    `json:"co"`
	Kode_pre_order        string `json:"kode_pre_order"`
	Tanggal               string `json:"tanggal"`
	Kode_nota             string `json:"kode_nota"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Kode_supplier         string `json:"kode_supplier"`
	Kode_gudang           string `json:"kode_gudang"`
	Status                int    `json:"status"`
}

type Input_Barang_Pre_Order_Request struct {
	Kode_stock         string `json:"kode_stock"`
	Tanggal_kadalurasa string `json:"tanggal_kadalurasa"`
	Jumlah_barang      string `json:"jumlah_barang"`
	Harga_pokok        string `json:"harga_pokok"`
}

type Input_Barang_Pre_Order_V2_Request struct {
	Co                    int     `json:"co"`
	Kode_barang_pre_order string  `json:"kode_barang_pre_order"`
	Kode_pre_order        string  `json:"kode_pre_order"`
	Kode_stock            string  `json:"kode_stock"`
	Tanggal_kadaluarsa    string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang         float64 `json:"jumlah_barang"`
	Harga                 int64   `json:"harga"`
	Total_harga           int64   `json:"total_harga"`
}

type Read_Pre_Order_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Pre_Order_Filter_Request struct {
	Tanggal_1     string `json:"tanggal_1"`
	Tanggal_2     string `json:"tanggal_2"`
	Kode_supplier string `json:"kode_supplier"`
}

type Update_Pre_Order_Kode_Request struct {
	Kode_barang_pre_order string `json:"kode_barang_pre_order"`
}

type Update_Pre_order_Request struct {
	Tanggal_kadaluarsa string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang      float64 `json:"jumlah_barang"`
	Harga              int64   `json:"harga"`
	Total_harga        int64   `json:"total_harga"`
}

type Kode_Pre_Order_Request struct {
	Kode_pre_order string `json:"kode_pre_order"`
}

type Update_Status_Pre_Order_Request struct {
	Status int `json:"status"`
}

type Move_Barang_Pre_Order_Request struct {
	Kode_stock         string  `json:"kode_stock"`
	Tanggal_kadaluarsa string  `json:"tanggal_kadaluarsa"`
	Jumlah_barang      float64 `json:"jumlah_barang"`
	Harga              int64   `json:"harga"`
}
