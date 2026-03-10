package handler

import (
	"net/http"
	"strings"

	"github.com/12Ndraaa/restapi-jualbeli/service"
)

type LaporanHandler struct {
	svc service.LaporanService
}

func NewLaporanHandler(svc service.LaporanService) *LaporanHandler {
	return &LaporanHandler{svc: svc}
}

func (h *LaporanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/laporan")
	path = strings.Trim(path, "/")

	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method ga diizinkan",
		})
		return
	}

	switch path {
	case "faktur":
		h.RekapPerFaktur(w, r)
	case "item":
		h.RekapPerItem(w, r)
	case "stok":
		h.LihatStok(w, r)
	default:
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "endpoint tidak ditemukan",
		})
	}
}

func (h *LaporanHandler) RekapPerFaktur(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.RekapPerFaktur()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *LaporanHandler) RekapPerItem(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.RekapPerItem()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *LaporanHandler) LihatStok(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.LihatStok()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}
