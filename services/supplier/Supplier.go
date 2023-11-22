package supplier

import (
	"POS-BACKEND/db"
	"POS-BACKEND/models/request"
	"POS-BACKEND/models/response"
	"POS-BACKEND/tools"
	"net/http"
	"strconv"
)

func Input_Supplier(Request request.Input_Supplier_Request, Request_Barang request.Input_Barang_Supplier_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("supplier")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_supplier = "SP-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_supplier", "nama_supplier", "nomor_telpon", "kode_gudang").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	con_brg := db.CreateConGorm().Table("barang_supplier")

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_stock)
	Request_Barang.Kode_supplier = Request.Kode_supplier

	for i := 0; i < len(kode_stock); i++ {
		var Request_Barang_Input request.Input_Barang_Supplier_Request

		co = 0

		err = con_brg.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_Barang_Input.Co = co + 1
		Request_Barang_Input.Kode_barang_supplier = "SPB-" + strconv.Itoa(Request_Barang_Input.Co)
		Request_Barang_Input.Kode_supplier = Request.Kode_supplier
		Request_Barang_Input.Kode_stock = kode_stock[i]

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = co
			return res, err.Error
		}

		err = con_brg.Select("co", "kode_barang_supplier", "kode_supplier", "kode_stock").Create(&Request_Barang_Input)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

	}

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}

func Read_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Supplier_Response
	var obj_data response.Read_Supplier_Response

	con := db.CreateConGorm().Table("supplier")

	rows, err := con.Select("kode_supplier", "nama_supplier", "nomor_telpon").Where("kode_gudang = ?", Request.Kode_gudang).Rows()

	defer rows.Close()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	for rows.Next() {
		con_barang := db.CreateConGorm().Table("barang_supplier")
		var detail_data []response.Read_Barang_Supplier_Response
		rows.Scan(&obj_data.Kode_supplier, &obj_data.Nama_supplier, &obj_data.Nomor_telpon)

		err := con_barang.Select("barang_supplier.kode_stock", "nama_barang").Joins("join stock on barang_supplier.kode_stock = stock.kode_stock").Where("kode_supplier = ?", obj_data.Kode_supplier).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		obj_data.Barang_supplier = detail_data

		data = append(data, obj_data)
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}

	return res, nil
}

func Dropdown_Nama_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {

	var res response.Response
	var Nama_supplier []response.Read_Dropdown_Nama_Supplier_Response

	con := db.CreateConGorm().Table("supplier")

	err := con.Select("kode_supplier", "nama_supplier").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&Nama_supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_supplier
		return res, err
	}

	if Nama_supplier == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Nama_supplier

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Nama_supplier
	}

	return res, nil
}

func Delete_Supplier(Request request.Delete_Supplier_Request) (response.Response, error) {
	var res response.Response

	var supplier []string

	con_masuk := db.CreateConGorm().Table("stock_masuk")

	err := con_masuk.Select("kode_supplier").Where("kode_supplier = ?", Request.Kode_supplier).Scan(&supplier).Error

	if supplier == nil && err == nil {
		con := db.CreateConGorm().Table("satuan_barang")

		err := con.Where("kode_satuan_barang=?", Request.Kode_supplier).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Erorr karena ada condition yang tidak terpenuhi"
		res.Data = Request
		return res, err
	}

	return res, nil
}

func Dropdown_Barang_Supplier(Request request.Read_Barang_Supplier_Request) (response.Response, error) {
	var res response.Response
	var Barang_Supplier []response.Read_Barang_Supplier_Response

	con := db.CreateConGorm().Table("barang_supplier")

	err := con.Select("barang_supplier.kode_stock", "nama_barang").Joins("JOIN stock s on s.kode_stock = barang_supplier.kode_stock").Where("kode_supplier = ?", Request.Kode_supplier).Scan(&Barang_Supplier).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang_Supplier
		return res, err
	}

	if Barang_Supplier == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Barang_Supplier

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = Barang_Supplier
	}

	return res, nil
}