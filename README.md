# Backend Task

## Introduction

This project provides  backend service for a flight tracking application that provides real-time flight
information using ADS-B data.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Go Installation](#go-installation)
- [Project Setup](#project-setup)
- [Directory Structure](#directory-structure)
- [Basic Authentication](#basic-authentication)
- [API Routes](#api-routes)


## Prerequisites

Go language: This project is written in Go and requires a working Go environment to run.

## Go Installation

If you don't have Go installed, follow the official installation instructions for your operating system: https://go.dev/doc/install

## Project Setup

Clone the repository and install dependencies:
```bash
git clone https://github.com/riteshinheritx/backend-take-home
cd backend-take-home
go mod tidy
```

## Directory Structure
```bash
/backend-take-home  
  ├──/pkg                  
        ├──flight-handler   # For business logic 
  ├──.air.toml              # For hot reloading
  ├──.env                   # For all secret/environment variables
  ├──.gitignore
  ├── flight_data.json      # Act as a database
  ├── main.go               # Application entry point
  └── go.mod                # Go modules file
  ├── README.md             # For documentation
```

* Right now, we didn't include .env in gitignore file for task submission

## Basic Authentication

This project implements basic authentication. All API requests must be authenticated with the following credentials to be successful:

Username: admin  
Password : admin@123

* For running it in postman you need to select Basic Auth in Authorization and then paste this credentials. 

## API Routes

### - Home Route [GET] ("/)
This route returns a simple response, typically used for health checks or basic information about the application.

### - Flights Data Route [POST] ("/getFlightData")

This route allows authorized users to retrieve all flight data with pagination and limit.

* For pagination query param : "page = any number"
* For limit query param      : "limit = any number"

### - Search Flight Information [POST] ("/searchFlightInfo")

This route allows authorized users to search for flight information based on specific criteria.

* For this route this kind of payload is required:-
```bash
{
    "flight_icao": "0E154",
    "flight_iata": "B354"
}
```

## Feedback

Thank for reading the documentation. We welcome all kinds of feedback on this project! Please feel free to create issues or pull requests on the project's GitHub repository.