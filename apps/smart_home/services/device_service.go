package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/practicumstudent2025/architecture-warmhouse/apps/smart_home/models"
)

var (
	ErrDeviceNotFound     = errors.New("device not found")
	ErrInvalidParameters  = errors.New("invalid parameters")
	ErrDuplicateDevice   = errors.New("device with this serial number already exists")
)

// DeviceRepository представляет интерфейс для работы с хранилищем устройств
type DeviceRepository interface {
	Create(ctx context.Context, device *models.Device) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Device, error)
	FindBySerialNumber(ctx context.Context, serialNumber string) (*models.Device, error)
}

// CommandRepository представляет интерфейс для работы с хранилищем команд
type CommandRepository interface {
	Create(ctx context.Context, command *models.DeviceCommand) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.DeviceCommand, error)
}

// DeviceService представляет сервис для работы с устройствами
type DeviceService struct {
	deviceRepo  DeviceRepository
	commandRepo CommandRepository
}

// NewDeviceService создает новый экземпляр DeviceService
func NewDeviceService(deviceRepo DeviceRepository, commandRepo CommandRepository) *DeviceService {
	return &DeviceService{
		deviceRepo:  deviceRepo,
		commandRepo: commandRepo,
	}
}

// RegisterDevice регистрирует новое устройство
func (s *DeviceService) RegisterDevice(ctx context.Context, req *models.RegisterDeviceRequest) (*models.RegisterDeviceResponse, error) {
	// Проверяем, не существует ли уже устройство с таким серийным номером
	existing, err := s.deviceRepo.FindBySerialNumber(ctx, req.SerialNumber)
	if err != nil && !errors.Is(err, ErrDeviceNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, ErrDuplicateDevice
	}

	device := &models.Device{
		ID:              uuid.New(),
		SerialNumber:    req.SerialNumber,
		TypeID:          req.TypeID,
		HouseID:         req.HouseID,
		Status:          models.DeviceStatusActive,
		FirmwareVersion: "1.2.3",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := s.deviceRepo.Create(ctx, device); err != nil {
		return nil, err
	}

	return &models.RegisterDeviceResponse{
		DeviceID:        device.ID,
		Status:          device.Status,
		FirmwareVersion: device.FirmwareVersion,
	}, nil
}

// SendCommand отправляет команду на устройство
func (s *DeviceService) SendCommand(ctx context.Context, deviceID uuid.UUID, req *models.SendCommandRequest) (*models.SendCommandResponse, error) {
	// Проверяем существование устройства
	device, err := s.deviceRepo.FindByID(ctx, deviceID)
	if err != nil {
		if errors.Is(err, ErrDeviceNotFound) {
			return nil, ErrDeviceNotFound
		}
		return nil, err
	}

	// Проверяем статус устройства
	if device.Status != models.DeviceStatusActive {
		return nil, errors.New("device is not active")
	}

	command := &models.DeviceCommand{
		ID:         uuid.New(),
		DeviceID:   deviceID,
		Type:       req.CommandType,
		Parameters: req.Parameters,
		Status:     models.CommandStatusQueued,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.commandRepo.Create(ctx, command); err != nil {
		return nil, err
	}

	return &models.SendCommandResponse{
		CommandID: command.ID,
		Status:    command.Status,
	}, nil
} 