package model

import "time"

type RekapPerFaktur struct {
	No        string       `json:"trxno"`
	Tanggal   time.Time `json:"tgl"`
	PartnerID int       `json:"bp_id"`
	Total     float64   `json:"total"`
}

type RekapPerItem struct {
	NamaItem string  `json:"nama_item"`
	Qty      int     `json:"qty"`
	Total    float64 `json:"total"`
}

type StokInfo struct {
	Nama string `json:"nama"`
	NamaItem string  `json:"nama_item"`
	Qty      int     `json:"qty"`
}
