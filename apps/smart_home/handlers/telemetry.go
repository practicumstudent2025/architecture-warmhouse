package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/models"
	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/services"
)

// TelemetryHandler HTTP-обработчик для телеметрии
type TelemetryHandler struct {
	telemetryService *services.TelemetryService
}

func NewTelemetryHandler(telemetryService *services.TelemetryService) *TelemetryHandler {
	return &TelemetryHandler{telemetryService: telemetryService}
}

func (h *TelemetryHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var telemetry models.DeviceTelemetry
	if err := json.NewDecoder(r.Body).Decode(&telemetry); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid request body"}`))
		return
	}
	if err := h.telemetryService.HandleTelemetry(r.Context(), &telemetry); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal error"}`))
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
