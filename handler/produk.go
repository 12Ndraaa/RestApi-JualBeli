package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

type ProdukHandler struct {
	service service.ProdukService
}

func NewProdukHandler(service service.ProdukService) *ProdukHandler {
	return &ProdukHandler{service: service}
}

// Helper ngirim response json
func writeJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// Router ngarahin request ke method yg bener
func (h *ProdukHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/produk")
	path = strings.Trim(path, "/")

	if path == "" {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, "method ga diizinkan", nil)
		}
	} else {
		id, err := strconv.Atoi(path)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, "id ga valid", nil)
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r, id)
		case http.MethodPut:
			h.Update(w, r, id)
		case http.MethodDelete:
			h.Delete(w, r, id)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, "method ga diizinkan", nil)
		}
	}
}

// respon semua
func (h *ProdukHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if len(data) == 0 {
		writeJSON(w,http.StatusOK, "gaada produk samsek", nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", data)
}

// respon per id
func (h *ProdukHandler) GetById(w http.ResponseWriter, r *http.Request, id int) {
	data, err := h.service.GetById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", data)
}

// parse & simpan
func (h *ProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, "request body ga valid", nil)
		return
	}
	data, err := h.service.Create(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusCreated, "produk berhasil dibuat", data)
}

// parse & ubah
func (h *ProdukHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, "request body ga valid", nil)
		return
	}
	p.ID = id

	data, err := h.service.Update(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "produk berhasil diubah", data)
}

// panggil hapus
func (h *ProdukHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "produk berhasil dihapus", nil)
}