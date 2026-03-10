package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/12Ndraaa/restapi-jualbeli/config"
	"github.com/12Ndraaa/restapi-jualbeli/handler"
	"github.com/12Ndraaa/restapi-jualbeli/repository"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

func main() {
	//koneksi ke db dlu
	db := config.InitDB()
	defer db.Close()

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

	// jlnin srver
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Printf("Server Jalan di Port: %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
