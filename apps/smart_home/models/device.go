package models

import (
	"time"

	"github.com/google/uuid"
)

// DeviceStatus представляет статус устройства
type DeviceStatus string

const (
	DeviceStatusActive   DeviceStatus = "ACTIVE"
	DeviceStatusInactive DeviceStatus = "INACTIVE"
	DeviceStatusPending  DeviceStatus = "PENDING"
)

// Device представляет IoT устройство
type Device struct {
	ID              uuid.UUID     `json:"id" db:"id"`
	SerialNumber    string        `json:"serialNumber" db:"serial_number"`
	TypeID          uuid.UUID     `json:"typeId" db:"type_id"`
	HouseID         uuid.UUID     `json:"houseId" db:"house_id"`
	Status          DeviceStatus  `json:"status" db:"status"`
	FirmwareVersion string        `json:"firmwareVersion" db:"firmware_version"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time     `json:"updatedAt" db:"updated_at"`
}

// RegisterDeviceRequest представляет запрос на регистрацию устройства
type RegisterDeviceRequest struct {
	SerialNumber string    `json:"serialNumber" validate:"required"`
	TypeID       uuid.UUID `json:"typeId" validate:"required"`
	HouseID      uuid.UUID `json:"houseId" validate:"required"`
}

// RegisterDeviceResponse представляет ответ на регистрацию устройства
type RegisterDeviceResponse struct {
	DeviceID        uuid.UUID    `json:"deviceId"`
	Status          DeviceStatus `json:"status"`
	FirmwareVersion string       `json:"firmwareVersion"`
}

// CommandStatus представляет статус команды
type CommandStatus string

const (
	CommandStatusQueued    CommandStatus = "QUEUED"
	CommandStatusSent      CommandStatus = "SENT"
	CommandStatusDelivered CommandStatus = "DELIVERED"
	CommandStatusExecuted  CommandStatus = "EXECUTED"
	CommandStatusFailed    CommandStatus = "FAILED"
)

// CommandType представляет тип команды
type CommandType string

const (
	CommandTypeSetTemperature CommandType = "SET_TEMPERATURE"
	CommandTypeSetMode        CommandType = "SET_MODE"
	CommandTypeTurnOn        CommandType = "TURN_ON"
	CommandTypeTurnOff       CommandType = "TURN_OFF"
)

// DeviceCommand представляет команду для устройства
type DeviceCommand struct {
	ID         uuid.UUID     `json:"id" db:"id"`
	DeviceID   uuid.UUID     `json:"deviceId" db:"device_id"`
	Type       CommandType   `json:"commandType" db:"command_type"`
	Parameters interface{}   `json:"parameters" db:"parameters"`
	Status     CommandStatus `json:"status" db:"status"`
	CreatedAt  time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time     `json:"updatedAt" db:"updated_at"`
}

// SendCommandRequest представляет запрос на отправку команды
type SendCommandRequest struct {
	CommandType CommandType `json:"commandType" validate:"required"`
	Parameters  interface{} `json:"parameters" validate:"required"`
}

// SendCommandResponse представляет ответ на отправку команды
type SendCommandResponse struct {
	CommandID uuid.UUID     `json:"commandId"`
	Status    CommandStatus `json:"status"`
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
} 