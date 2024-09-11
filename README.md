# Go Reference API

## Overview

This API is a simple reference for a minimalistic REST API written in Go. It utilizes the built-in
http.ServeMux to handle routing and the built-in http package to handle requests and responses.

## Installation

To install the API, clone the repository and run the following command:

```bash
go install
```

## Directory Structure

```text
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── handlers.go
│   └── health/
│       └── health.go
├── pkg/
│   ├── adapters/
│   │   └── database.go
│   ├── middleware/
│   │   └── auth.go
│   └── mylibrary/
│       └── mylibrary.go
├── swagger-ui/
│   └── (Swagger UI files)
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## Routes

The API has the following routes:

- `/` - Home route
- `/health` - Health check route


Sending a GET request to the /healthz route

```shell
curl -X GET http://localhost:8080/healthz
```

Sned a GET request to the /readyz route

```shell
curl -X GET http://localhost:8080/readyz
```
