package kartu_stock

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Read_Kartu_Stock(Request request.Read_Kartu_Stock_Request) (response.Response, error) {
	var res response.Response
	var read_sup []response.Read_Sup_Plus_Barang_Response

	if Request.Tanggal_1 == "" && Request.Tanggal_2 == "" {

		date_max_time := time.Now()
		Request.Tanggal_2 = date_max_time.Format("2006-01-02")

		date_min := date_max_time.Format("2006-01")
		Request.Tanggal_1 = date_min + "-01"

		fmt.Println(Request.Tanggal_1)
		fmt.Println(Request.Tanggal_2)

	} else {
		date_1, _ := time.Parse("02-01-2006", Request.Tanggal_1)
		Request.Tanggal_1 = date_1.Format("2006-01-02")

		date_2, _ := time.Parse("02-01-2006", Request.Tanggal_2)
		Request.Tanggal_2 = date_2.Format("2006-01-02")
	}

	if Request.Kode_supplier != "" && Request.Kode_stock != "" {

		con_supplier := db.CreateConGorm().Table("supplier")

		err := con_supplier.Select("supplier.kode_supplier", "nama_supplier", "bs.kode_stock", "nama_barang").Joins("JOIN barang_supplier bs ON bs.kode_supplier = supplier.kode_supplier").Joins("JOIN stock s ON s.kode_stock =  bs.kode_stock").Where("supplier.kode_gudang = ? && bs.kode_stock = ? && supplier.kode_supplier = ?", Request.Kode_gudang, Request.Kode_stock, Request.Kode_supplier).Scan(&read_sup)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		fmt.Println(read_sup)

	} else if Request.Kode_supplier != "" {

		con_supplier := db.CreateConGorm().Table("supplier")

		err := con_supplier.Select("supplier.kode_supplier", "nama_supplier", "bs.kode_stock", "nama_barang").Joins("JOIN barang_supplier bs ON bs.kode_supplier = supplier.kode_supplier").Joins("JOIN stock s ON s.kode_stock =  bs.kode_stock").Where("supplier.kode_gudang = ? && supplier.kode_supplier = ?", Request.Kode_gudang, Request.Kode_supplier).Scan(&read_sup)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		fmt.Println(read_sup)

	} else if Request.Kode_stock != "" {

		con_supplier := db.CreateConGorm().Table("supplier")

		err := con_supplier.Select("supplier.kode_supplier", "nama_supplier", "bs.kode_stock", "nama_barang").Joins("JOIN barang_supplier bs ON bs.kode_supplier = supplier.kode_supplier").Joins("JOIN stock s ON s.kode_stock =  bs.kode_stock").Where("supplier.kode_gudang = ? && bs.kode_stock = ?", Request.Kode_gudang, Request.Kode_stock).Scan(&read_sup)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		fmt.Println(read_sup)

	} else {

		con_supplier := db.CreateConGorm().Table("supplier")

		err := con_supplier.Select("supplier.kode_supplier", "nama_supplier", "bs.kode_stock", "nama_barang").Joins("JOIN barang_supplier bs ON bs.kode_supplier = supplier.kode_supplier").Joins("JOIN stock s ON s.kode_stock =  bs.kode_stock").Where("supplier.kode_gudang = ?", Request.Kode_gudang).Scan(&read_sup)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

	}

	var data response.Kartu_Stock_Response
	var read_kartu_stock []response.Kartu_Stock_Response
	var raw_data []response.Read_Raw_Kartu_Stock_Response

	con_stock_raw := db.CreateConGorm().Table("stock_keluar_masuk")

	err := con_stock_raw.Select("kode_stock_keluar_masuk", "tanggal", "kode").Where("kode_gudang = ? && (tanggal >=? && tanggal <=?)", Request.Kode_gudang, Request.Tanggal_1, Request.Tanggal_2).Order("tanggal ASC").Scan(&raw_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for j := 0; j < len(read_sup); j++ {
		sisa := 0.0
		total_masuk := 0.0
		total_keluar := 0.0

		var detail []response.Detail_Kartu_Stock_Response

		data.Nama_barang = read_sup[j].Nama_barang
		data.Nama_supplier = read_sup[j].Nama_supplier

		for i := 0; i < len(raw_data); i++ {
			var detail_raw []response.Detail_Kartu_Stock_Response

			if strings.HasPrefix(raw_data[i].Kode_stock_keluar_masuk, "SM") {
				con_stock_raw := db.CreateConGorm().Table("stock_keluar_masuk")
				err = con_stock_raw.Select("DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "jumlah_barang").Joins("JOIN barang_stock_keluar_masuk bskm ON stock_keluar_masuk.kode_stock_keluar_masuk = bskm.kode_stock_keluar_masuk").Where("stock_keluar_masuk.kode_stock_keluar_masuk = ? && kode =? && kode_stock = ? ", raw_data[i].Kode_stock_keluar_masuk, read_sup[j].Kode_supplier, read_sup[j].Kode_stock).Order("tanggal ASC").Scan(&detail_raw)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				for x := 0; x < len(detail_raw); x++ {
					detail_raw[x].Keterangan = "MASUK"
					total_masuk = total_masuk + detail_raw[x].Jumlah_barang
					sisa = sisa + detail_raw[x].Jumlah_barang
					detail_raw[x].Sisa = sisa
				}

			} else if strings.HasPrefix(raw_data[i].Kode_stock_keluar_masuk, "SK") {

				con_stock_raw := db.CreateConGorm().Table("pengurangan_stock")
				err = con_stock_raw.Select("jumlah_barang").Joins("JOIN barang_stock_keluar_masuk bskm ON pengurangan_stock.kode_barang_keluar = bskm.kode_barang_keluar_masuk").Where("kode_stock_keluar = ? && kode_supplier = ? && kode_stock = ? ", raw_data[i].Kode_stock_keluar_masuk, read_sup[j].Kode_supplier, read_sup[j].Kode_stock).Scan(&detail_raw)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				for x := 0; x < len(detail_raw); x++ {
					date, _ := time.Parse("2006-01-02", raw_data[i].Tanggal)
					detail_raw[x].Tanggal = date.Format("02-01-2006")
					detail_raw[x].Keterangan = "KELUAR"
					total_keluar = total_keluar + detail_raw[x].Jumlah_barang
					sisa = sisa - detail_raw[x].Jumlah_barang
					detail_raw[x].Sisa = sisa
				}

			} else if strings.HasPrefix(raw_data[i].Kode_stock_keluar_masuk, "AU") {
				var detail_raw_single response.Detail_Kartu_Stock_Response
				var audit_raw []response.Audit_RAW_response
				con_stock_raw := db.CreateConGorm().Table("audit")

				err = con_stock_raw.Select("stock_dalam_sistem, stock_rill").Joins("JOIN detail_audit da ON da.kode_audit=audit.kode_audit").Where("da.`status`= 0 AND audit.kode_audit = ? AND kode_supplier = ? AND kode_stock = ? AND audit.`status` = 1", raw_data[i].Kode_stock_keluar_masuk, read_sup[j].Kode_supplier, read_sup[j].Kode_stock).Scan(&audit_raw)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				selisih := 0.0
				for y := 0; y < len(audit_raw); y++ {
					selisih = selisih + (audit_raw[y].Stock_rill - audit_raw[y].Stock_dalam_sistem)
				}

				date, _ := time.Parse("2006-01-02", raw_data[i].Tanggal)
				detail_raw_single.Tanggal = date.Format("02-01-2006")
				detail_raw_single.Jumlah_barang = selisih

				if selisih > 0.0 {
					detail_raw_single.Keterangan = "AUDIT MASUK"
					total_masuk = total_masuk + detail_raw_single.Jumlah_barang
					sisa = sisa + detail_raw_single.Jumlah_barang

				} else if selisih == 0.0 {

					detail_raw_single.Keterangan = "AUDIT"
					total_masuk = total_masuk + detail_raw_single.Jumlah_barang
					sisa = sisa - detail_raw_single.Jumlah_barang

				} else if selisih < 0.0 {

					detail_raw_single.Keterangan = "AUDIT KELUAR"
					total_keluar = total_masuk + detail_raw_single.Jumlah_barang
					sisa = sisa - detail_raw_single.Jumlah_barang

				}

				detail_raw_single.Sisa = sisa

				detail_raw = append(detail_raw, detail_raw_single)
			}

			if len(detail_raw) > 0 {
				detail = append(detail, detail_raw...)
			}
		}

		data.Jumlah_stock_masuk = total_masuk
		data.Jumlah_stock_keluar = total_keluar
		data.Detail_kartu_stock = detail

		if len(detail) > 0 {
			read_kartu_stock = append(read_kartu_stock, data)
		}

	}

	if read_kartu_stock == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = read_kartu_stock

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = read_kartu_stock
	}

	return res, nil
}
