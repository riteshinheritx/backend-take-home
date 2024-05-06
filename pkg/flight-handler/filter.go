package flighthandler

import (
	"reflect"
)

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// filterFlights Function to filter flight data based on parameters
func filterFlights(flights FlightData, params map[string]string) FlightData {

	// If no parameters are provided, return all flights
	if len(params) == 0 {
		return flights
	}

	// Apply filters
	filteredFlights := make(FlightData, 0)

	for _, flight := range flights {
		if filterWithParams(flight, params) {
			filteredFlights = append(filteredFlights, flight)
		}
	}

	return filteredFlights
}

// filterWithParams filters Flight struct based on the provided parameters
func filterWithParams(flight Flight, params map[string]string) bool {
	// Iterate over each parameter
	for key, value := range params {
		// Get the corresponding struct field name using the flightFieldKeyMap
		fieldKeyName := flightFieldKeyMap[key]

		// Retrieve the field from the Flight struct by name
		field := reflect.ValueOf(flight).FieldByName(fieldKeyName)
		// Check if the field is valid
		if field.IsValid() {
			// If the field is a valid string field, check if its value matches the provided value
			if fieldValue := field.String(); value != "" && fieldValue != value {
				return false // If it doesn't match, return false
			}
		}
	}
	return true // If all parameters match, return true
}

// Function to apply pagination to filtered flights
func paginateFlights(filteredFlights []Flight, pagination PaginationParams) []Flight {
	// Adjust page to start from 0-based index
	page := pagination.Page - 1
	startIndex := page * pagination.PageSize
	endIndex := startIndex + pagination.PageSize

	// Check if start index is out of bounds
	if startIndex >= len(filteredFlights) {
		return nil
	}

	// Check if end index exceeds the length of filtered flights
	if endIndex > len(filteredFlights) {
		endIndex = len(filteredFlights)
	}

	return filteredFlights[startIndex:endIndex]
}
