from flask import Flask, request, jsonify
import random

app = Flask(__name__)

SENSOR_LOCATIONS = {
    "1": "Living Room",
    "2": "Bedroom",
    "3": "Kitchen"
}
LOCATION_SENSORS = {v: k for k, v in SENSOR_LOCATIONS.items()}

@app.route('/temperature')
def get_temperature():
    location = request.args.get('location', '')
    sensor_id = request.args.get('sensorId', '')

    # If no location is provided, use a default based on sensor ID
    if not location:
        location = SENSOR_LOCATIONS.get(sensor_id, 'Unknown')

    # If no sensor ID is provided, generate one based on location
    if not sensor_id:
        sensor_id = LOCATION_SENSORS.get(location, '0')

    temperature = round(random.uniform(15.0, 30.0), 2)
    return jsonify({
        'sensorId': sensor_id,
        'location': location,
        'temperature': temperature
    })

@app.route('/temperature/<sensor_id>')
def get_temperature_by_id(sensor_id):
    location = SENSOR_LOCATIONS.get(sensor_id, 'Unknown')
    temperature = round(random.uniform(15.0, 30.0), 2)
    return jsonify({
        'sensorId': sensor_id,
        'location': location,
        'temperature': temperature
    })

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081) 