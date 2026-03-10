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
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Router ngarahin request ke method yg bener
func (h *ProdukHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// url /produk atau /produk/1
	path := strings.TrimPrefix(r.URL.Path, "/produk")
	path = strings.Trim(path, "/")

	if path == "" {
		// /produk
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "method ga diizinkan",
			})
		}
	} else {
		// /produk/1
		id, err := strconv.Atoi(path)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "id ga valid",
			})
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
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "method ga diizinkan",
			})
		}
	}
}

// respon semua
func (h *ProdukHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, data)
}

// respon per id
func (h *ProdukHandler) GetById(w http.ResponseWriter, r *http.Request, id int) {
	data, err := h.service.GetById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, data)
}

// parse & simpan
func (h *ProdukHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "request body ga valid",
		})
		return
	}
	data, err := h.service.Create(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusCreated, data)
}

// parse & ubah
func (h *ProdukHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "request body ga valid",
		})
		return
	}
	p.ID = id // jd yg dipke id dri url

	data, err := h.service.Update(&p)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, data)
}

// panggil hapus
func (h *ProdukHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "produk sukses dihapus",
	})
}
