package routes

import (
	"database/sql"
	"net/http"

	"github.com/12Ndraaa/restapi-jualbeli/handler"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

func RegisterRoutes(db *sql.DB) {

	// buat repo
	produkRepo := repository.NewProdukRepository(db)
	mitrabisnisRepo := repository.NewMitraBisnisRepository(db)
	gudangRepo := repository.NewGudangRepository(db)
	pembelianRepo := repository.NewPembelianRepository(db)
	penjualanRepo := repository.NewPenjualanRepository(db)
	laporanRepo := repository.NewLaporanRepository(db)

	// buat service
	produkService := service.NewProdukService(produkRepo)
	mitrabisnisService := service.NewMitraBisnisService(mitrabisnisRepo)
	gudangService := service.NewGudangService(gudangRepo)
	pembelianSvc := service.NewPembelianService(pembelianRepo)
	penjualanService := service.NewPenjualanService(penjualanRepo)
	laporanService := service.NewLaporanService(laporanRepo)

	// buat handler
	produkHandler := handler.NewProdukHandler(produkService)
	mitrabisnisHandler := handler.NewMitraBisnisHandler(mitrabisnisService)
	pembelianHandler := handler.NewPembelianHandler(pembelianSvc)
	gudangHandler := handler.NewGudangHandler(gudangService)
	penjualanHandler := handler.NewPenjualanHandler(penjualanService)
	laporanHandler := handler.NewLaporanHandler(laporanService)

	// daftarin route dlu
	http.Handle("/produk", produkHandler)        // base produk
	http.Handle("/produk/", produkHandler)       // slash
	http.Handle("/mitra", mitrabisnisHandler)    // base mitra
	http.Handle("/mitra/", mitrabisnisHandler)   // slash
	http.Handle("/gudang", gudangHandler)        // base gudang
	http.Handle("/gudang/", gudangHandler)       // slash
	http.Handle("/pembelian", pembelianHandler)  // base pembelian
	http.Handle("/pembelian/", pembelianHandler) // slash
	http.Handle("/penjualan", penjualanHandler)  // base penjualan
	http.Handle("/penjualan/", penjualanHandler) // slash
	http.Handle("/laporan/", laporanHandler)     // semua laporan
}
