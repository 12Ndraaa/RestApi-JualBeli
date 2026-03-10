package model

import "time"

type TrxPenjualan struct {
	ID        int       `json:"id"`
	No        string    `json:"trxno"`
	PartnerID int       `json:"bp_id"`
	Tanggal   time.Time `json:"tgl"`
	Diskon    float64   `json:"diskon"`
	SubTotal  float64   `json:"subtotal"`
	Total     float64   `json:"total"`
}

type TrxDetailPenjualan struct {
	SaleID int     `json:"sale_id"`
	DetailNum  int     `json:"dno"`
    GudangID   int     `json:"gudang_id"`
	ItemID     int     `json:"item_id"`
	Qty        int     `json:"qty"`
	Harga      float64 `json:"harga"`
	Diskon     float64 `json:"diskon"`
	SubTotal   float64 `json:"subtotal"`
}
