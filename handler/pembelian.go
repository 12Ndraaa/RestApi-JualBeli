package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

type PembelianHandler struct {
	svc service.PembelianService
}

func NewPembelianHandler(svc service.PembelianService) *PembelianHandler {
	return &PembelianHandler{svc: svc}
}

func (h *PembelianHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/pembelian")
	path = strings.Trim(path, "/")

	if path == "" {
		// /pembelian
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, "method ga diizinin", nil)
		}
	} else {
		// /pembelian/1
		id, err := strconv.Atoi(path)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, "id ga valid", nil)
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r, id)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, "method ga diizinin", nil)
		}
	}
}

func (h *PembelianHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if len(result) == 0 {
		writeJSON(w,http.StatusOK, "gaada pembelian samsek", nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", result)
}

func (h *PembelianHandler) GetById(w http.ResponseWriter, r *http.Request, id int) {
	result, err := h.svc.GetById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusOK, "success", result)
}

func (h *PembelianHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Header  model.TrxPembelian         `json:"header"`
		Details []model.TrxDetailPembelian `json:"details"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, "request ga valid", nil)
		return
	}

	result, err := h.svc.Create(&req.Header, req.Details)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	writeJSON(w, http.StatusCreated, "success", result)
}