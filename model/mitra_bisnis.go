package model

type MitraBisnis struct {
	ID int `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
	TipeMitra string `json:"tipe_mitra"`
}

