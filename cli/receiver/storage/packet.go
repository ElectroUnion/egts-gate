package storage

import (
	"encoding/json"
)

type NavRecord struct {
	ServerID            string         `json:"server_id"`
	IMEI                string         `json:"IMEI"`
	ClientID            uint32         `json:"client_id"`
	ReceivedDt          string         `json:"received_dt"`
	PacketID            uint32         `json:"packet_id"`
	NavigationTimestamp int64          `json:"navigation_unix_time"`
	AnSensors           []AnSensor     `json:"an_sensors"`
	LiquidSensors       []LiquidSensor `json:"liquid_sensors"`
	LocStates           []LocState     `json:"loc_states"`
	LocStatesCount      int32          `json:"loc_states_count"`
}

func (eep *NavRecord) ToBytes() ([]byte, error) {
	return json.Marshal(eep)
}

type LocState struct {
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lng"`
	Speed        uint16  `json:"speed"`
	NavigationDt string  `json:"navigation_dt"`
}

type LiquidSensor struct {
	SensorNumber uint8   `json:"sensor_number"`
	ModuleNumber uint16  `json:"module_number"`
	ErrorFlag    string  `json:"error_flag"`
	ValueUnit    string  `json:"value_unit"`
	ValueMm      uint32  `json:"value_mm"`
	ValueL       float64 `json:"value_l"`
}

type AnSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	Value        uint32 `json:"value"`
}
