package storage

import (
	"encoding/json"
)

type NavRecord struct {
	Client              uint32         `json:"client"`
	PacketID            uint32         `json:"packet_id"`
	IMEI                string         `json:"IMEI"`
	NavigationTimestamp int64          `json:"navigation_unix_time"`
	ReceivedTimestamp   int64          `json:"received_unix_time"`
	Pdop                uint16         `json:"pdop"`
	Hdop                uint16         `json:"hdop"`
	Vdop                uint16         `json:"vdop"`
	Nsat                uint8          `json:"nsat"`
	Ns                  uint16         `json:"ns"`
	AnSensors           []AnSensor     `json:"an_sensors"`
	LiquidSensors       []LiquidSensor `json:"liquid_sensors"`
	LocStates           []LocState     `json:"loc_states"`
	LocStatesCount      int32          `json:"loc_states_count"`
}

func (eep *NavRecord) ToBytes() ([]byte, error) {
	return json.Marshal(eep)
}

type LocState struct {
	Latitude            float64 `json:"lat"`
	Longitude           float64 `json:"lng"`
	Speed               uint16  `json:"speed"`
	NavigationTimestamp int64   `json:"navigation_unix_time"`
}

type LiquidSensor struct {
	SensorNumber uint8   `json:"sensor_number"`
	ErrorFlag    string  `json:"error_flag"`
	ValueMm      uint32  `json:"value_mm"`
	ValueL       float64 `json:"value_l"`
}

type AnSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	Value        uint32 `json:"value"`
}
