package services

import (
	"context"
	"log"

	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/models"
)

// TelemetryService обрабатывает сообщения телеметрии устройств
type TelemetryService struct{}

func NewTelemetryService() *TelemetryService {
	return &TelemetryService{}
}

// HandleTelemetry обрабатывает входящее сообщение телеметрии
func (s *TelemetryService) HandleTelemetry(ctx context.Context, telemetry *models.DeviceTelemetry) error {
	log.Printf("Received telemetry: deviceId=%s timestamp=%s metrics=%+v", telemetry.DeviceID, telemetry.Timestamp, telemetry.Metrics)
	// Здесь может быть сохранение в БД, публикация в очередь и т.д.
	return nil
} 