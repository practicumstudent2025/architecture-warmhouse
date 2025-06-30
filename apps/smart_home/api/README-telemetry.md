# Telemetry Service

## AsyncAPI спецификация

Файл: `telemetry-asyncapi.yaml`

- Канал: `device.telemetry`
- Сообщение:
  - `deviceId` (uuid): идентификатор устройства
  - `timestamp` (date-time): время снятия показаний
  - `metrics` (array): список метрик (например, температура, влажность)

Пример сообщения:
```json
{
  "deviceId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "timestamp": "2023-10-05T14:30:00Z",
  "metrics": [
    { "name": "temperature", "value": 22.5 },
    { "name": "humidity", "value": 45 }
  ]
}
```

## HTTP endpoint для тестирования

Можно отправить POST-запрос на `/telemetry`:

```
curl -X POST http://localhost:8080/telemetry \
  -H "Content-Type: application/json" \
  -d '{
    "deviceId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "timestamp": "2023-10-05T14:30:00Z",
    "metrics": [
      { "name": "temperature", "value": 22.5 },
      { "name": "humidity", "value": 45 }
    ]
  }'
```

В реальной системе сообщения будут поступать через брокер сообщений (например, NATS, Kafka) по каналу `device.telemetry`. 