# Smart Home Monolith & Temperature API

## Temperature API

Сервис для генерации случайной температуры по комнате или sensorId.

### Запуск

```bash
cd apps
# Собрать и запустить все сервисы
docker-compose up --build
```

### Эндпоинт

`GET /temperature?location=Living%20Room` или `GET /temperature?sensorId=1`

Ответ:
```json
{
  "sensorId": "1",
  "location": "Living Room",
  "temperature": 23.45
}
```

- Если не указан location, он определяется по sensorId (1 — Living Room, 2 — Bedroom, 3 — Kitchen)
- Если не указан sensorId, он определяется по location
- Температура всегда случайная

## Smart Home Monolith

(см. инструкции ниже для работы с датчиками и БД)

## Prerequisites

- Docker and Docker Compose

## Getting Started

### Option 1: Using Docker Compose (Recommended)

The easiest way to start the application is to use Docker Compose:

```bash
./init.sh
```

This script will:

1. Build and start the PostgreSQL and application containers
2. Wait for the services to be ready
3. Display information about how to access the API

Alternatively, you can run Docker Compose directly:

```bash
docker-compose up -d
```

The API will be available at http://localhost:8080

### Option 2: Manual setup

If you prefer to run the application without Docker:

1. Start the PostgreSQL database:

```bash
docker-compose up -d postgres
```

2. Build and run the application:

```bash
go build -o smarthome
./smarthome
```

## API Testing

A Postman collection is provided for testing the API. Import the `smarthome-api.postman_collection.json` file into Postman to get started.

## API Endpoints

- `GET /health` - Health check
- `GET /api/v1/sensors` - Get all sensors
- `GET /api/v1/sensors/:id` - Get a specific sensor
- `POST /api/v1/sensors` - Create a new sensor
- `PUT /api/v1/sensors/:id` - Update a sensor
- `DELETE /api/v1/sensors/:id` - Delete a sensor
- `PATCH /api/v1/sensors/:id/value` - Update a sensor's value and status
