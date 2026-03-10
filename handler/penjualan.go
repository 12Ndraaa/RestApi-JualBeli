package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/12Ndraaa/restapi-jualbeli/model"
	"github.com/12Ndraaa/restapi-jualbeli/service"
)

type PenjualanHandler struct {
	svc service.PenjualanService
}

func NewPenjualanHandler(svc service.PenjualanService) *PenjualanHandler {
	return &PenjualanHandler{svc: svc}
}

func (h *PenjualanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/penjualan")
	path = strings.Trim(path, "/")

	if path == "" {
		// /pembelian
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
		// /pembelian/1
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
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "method ga diizinkan",
			})
		}
	}
}

func (h *PenjualanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *PenjualanHandler) GetById(w http.ResponseWriter, r *http.Request, id int) {
	result, err := h.svc.GetById(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *PenjualanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Header  model.TrxPenjualan         `json:"header"`
		Details []model.TrxDetailPenjualan `json:"details"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "request tidak valid"})
		return
	}

	result, err := h.svc.Create(&req.Header, req.Details)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, result)
}
