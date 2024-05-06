package flighthandler

import (
	_ "embed"
	"encoding/json"
)

type Flight struct {
	Hex          string `json:"hex"`
	RegNumber    string `json:"reg_number"`
	AircraftICAO string `json:"aircraft_icao"`
	Flag         string `json:"flag"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
	Alt          string `json:"alt"`
	Dir          string `json:"dir"`
	Speed        string `json:"speed"`
	VSpeed       string `json:"v_speed"`
	Squawk       string `json:"squawk"`
	AirlineICAO  string `json:"airline_icao"`
	AirlineIATA  string `json:"airline_iata"`
	Number       string `json:"flight_number"`
	ICAO         string `json:"flight_icao"`
	IATA         string `json:"flight_iata"`
	Duration     string `json:"duration"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
}

var flightFieldKeyMap = map[string]string{
	"hex":           "Hex",
	"reg_number":    "RegNumber",
	"aircraft_icao": "AircraftICAO",
	"flag":          "Flag",
	"lat":           "Lat",
	"lng":           "Lng",
	"alt":           "Alt",
	"dir":           "Dir",
	"speed":         "Speed",
	"v_speed":       "VSpeed",
	"squawk":        "Squawk",
	"airline_icao":  "AirlineICAO",
	"airline_iata":  "AirlineIATA",
	"flight_number": "Number",
	"flight_icao":   "ICAO",
	"flight_iata":   "IATA",
	"duration":      "Duration",
	"updated":       "Updated",
	"status":        "Status",
}

type FlightData []Flight

func UnmarshalFlightData(data []byte) FlightData {
	var flights FlightData

	if err := json.Unmarshal(data, &flights); err != nil {
		return make(FlightData, 0)
	}

	return flights
}
