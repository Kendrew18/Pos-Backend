package jenis_pembayaran

import (
	"POS-BACKEND/models/request_kasir"
	"POS-BACKEND/services/kasir/jenis_pembayaran"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputJenisPembayaran(c echo.Context) error {

	var Request request_kasir.Input_Jenis_Pembayaran_Request
	Request.Nama_jenis_pembayaran = c.FormValue("nama_jenis_pembayaran")
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := jenis_pembayaran.Input_Jenis_Pembayaran(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func ReadJenisPembayaran(c echo.Context) error {
	var Request request_kasir.Read_Jenis_Pembayaran_Request
	Request.Kode_kasir = c.FormValue("kode_kasir")

	result, err := jenis_pembayaran.Read_Jenis_Pembayaran(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}

func DeleteJenisPembayaran(c echo.Context) error {
	var Request request_kasir.Delete_Jenis_Pembayaran_Request
	Request.Kode_jenis_pembayaran = c.FormValue("kode_jenis_pembayaran")

	result, err := jenis_pembayaran.Delete_Jenis_Pembayaran(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)

}
