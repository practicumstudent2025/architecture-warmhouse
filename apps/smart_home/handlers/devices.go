package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/models"
	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/services"
)

// DeviceHandler обрабатывает HTTP запросы для управления устройствами
type DeviceHandler struct {
	deviceService *services.DeviceService
}

// NewDeviceHandler создает новый экземпляр DeviceHandler
func NewDeviceHandler(deviceService *services.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		deviceService: deviceService,
	}
}

// RegisterRoutes регистрирует маршруты для обработки запросов
func (h *DeviceHandler) RegisterRoutes(r chi.Router) {
	r.Post("/devices", h.RegisterDevice)
	r.Post("/devices/{deviceId}/commands", h.SendCommand)
}

// writeJSON отправляет JSON ответ
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError отправляет JSON ответ с ошибкой
func writeError(w http.ResponseWriter, status int, err error) {
	resp := models.ErrorResponse{
		Code:    http.StatusText(status),
		Message: err.Error(),
	}
	writeJSON(w, status, resp)
}

// RegisterDevice обрабатывает запрос на регистрацию устройства
func (h *DeviceHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	resp, err := h.deviceService.RegisterDevice(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrDuplicateDevice):
			writeError(w, http.StatusConflict, err)
		case errors.Is(err, services.ErrInvalidParameters):
			writeError(w, http.StatusBadRequest, err)
		default:
			writeError(w, http.StatusInternalServerError, errors.New("internal server error"))
		}
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}

// SendCommand обрабатывает запрос на отправку команды устройству
func (h *DeviceHandler) SendCommand(w http.ResponseWriter, r *http.Request) {
	deviceID, err := uuid.Parse(chi.URLParam(r, "deviceId"))
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid device ID"))
		return
	}

	var req models.SendCommandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	resp, err := h.deviceService.SendCommand(r.Context(), deviceID, &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrDeviceNotFound):
			writeError(w, http.StatusNotFound, err)
		case errors.Is(err, services.ErrInvalidParameters):
			writeError(w, http.StatusBadRequest, err)
		default:
			writeError(w, http.StatusInternalServerError, errors.New("internal server error"))
		}
		return
	}

	writeJSON(w, http.StatusAccepted, resp)
} 