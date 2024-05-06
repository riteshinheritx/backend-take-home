package flighthandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Base struct {
	FlightData FlightData
}

type Handlers interface {
	List(c echo.Context) error
	Info(c echo.Context) error
}

func NewFlight(data []byte) Handlers {
	return &Base{
		FlightData: UnmarshalFlightData(data),
	}
}

// List Handler for /getFlightData endpoint
func (b Base) List(c echo.Context) error {

	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	numPage, err := strconv.Atoi(page)
	if err != nil || numPage < 1 {
		numPage = 1
	}
	numPageSize, err := strconv.Atoi(limit)
	if err != nil || numPageSize < 1 {
		numPageSize = 20
	}

	pagination := PaginationParams{
		Page:     numPage,
		PageSize: numPageSize,
	}

	// Parse request body to extract parameters
	var params map[string]string
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Filter flights based on parameters
	filtered := filterFlights(b.FlightData, params)

	result := paginateFlights(filtered, pagination)

	// Return filtered flight data
	return c.JSON(http.StatusOK, echo.Map{
		"pagination":    pagination,
		"filteredCount": len(filtered),
		"flights":       result,
		"total":         len(b.FlightData), // total flights
	})
}

// Info Handler for /searchFlightInfo endpoint
func (b Base) Info(c echo.Context) error {
	// Parse request body to extract parameters
	var params map[string]string
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Check if flight_icao is missing in the request body
    if _, ok := params["flight_icao"]; !ok {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "ICAO number is missing"})
    }

	// Check if flight_iata is missing in the request body
    if _, ok := params["flight_iata"]; !ok {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "IATA number is missing"})
    }

	// Find flight info based on flight_icao or flight_iata code
	var foundFlight *Flight
	for _, flight := range b.FlightData {
		if flight.ICAO == params["flight_icao"] && flight.IATA == params["flight_iata"] {
			foundFlight = &flight
			break
		}
	}


	// Return flight info if found, otherwise return 404
	if foundFlight != nil {
		return c.JSON(http.StatusOK, foundFlight)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Flight not found"})
}


// ExemptionRoute Handler for /searchFlightInfo endpoint
func (b Base) ExemptionRoute(c echo.Context) error {
	// Parse request body to extract parameters
	
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	
}
