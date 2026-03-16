package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

type GudangHandler struct {
	service service.GudangService
}

func NewGudangHandler(service service.GudangService) *GudangHandler {
	return &GudangHandler{service: service}
}

// Router ngarahin request ke method yg bener
func (h *GudangHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// url /produk atau /produk/1
	path := strings.TrimPrefix(r.URL.Path, "/gudang")
	path = strings.Trim(path, "/")

	if path == "" {
		// /produk
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, "method ga diizinkan", nil)
		}
	} else {
		// /produk/1
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
func (h *GudangHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if len(data) == 0 {
		writeJSON(w,http.StatusOK, "gaada gudang samsek", nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", data)
}

// respon per id
func (h *GudangHandler) GetById(w http.ResponseWriter, r *http.Request, id int) {
	data, err := h.service.GetById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", data)
}

// parse & simpan
func (h *GudangHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Gudang
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	data, err := h.service.Create(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusCreated, "sukses dibuat", data)
}

// parse & ubah
func (h *GudangHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var p model.Gudang
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, "request body ga valid", nil)
		return
	}
	p.ID = id // jd yg dipke id dri url

	data, err := h.service.Update(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", data)
}

// panggil hapus
func (h *GudangHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "gudang berhasil dihapus", nil)
}
