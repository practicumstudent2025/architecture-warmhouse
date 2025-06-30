package models

import "time"

// TelemetryMetric описывает одну метрику устройства
type TelemetryMetric struct {
	Name  string   `json:"name"`
	Value float64  `json:"value"`
}

// DeviceTelemetry описывает сообщение телеметрии устройства
type DeviceTelemetry struct {
	DeviceID  string            `json:"deviceId"`
	Timestamp time.Time         `json:"timestamp"`
	Metrics   []TelemetryMetric `json:"metrics"`
} 