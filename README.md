# Puasa Sunnah / Sunnah Fasting API

<div align="center">

[![Better Uptime Badge](https://betteruptime.com/status-badges/v1/monitor/ds3l.svg)](https://betteruptime.com/?utm_source=status_badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/granitebps/puasa-sunnah-api)](https://goreportcard.com/report/github.com/granitebps/puasa-sunnah-api)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/granitebps/puasa-sunnah-api)
![GitHub](https://img.shields.io/github/license/granitebps/puasa-sunnah-api)

</div>

## Link
- [API](https://api.puasa-sunnah.granitebps.com)
- [API Docs](https://api.puasa-sunnah.granitebps.com/swagger)

## Framework
- [Fiber](https://gofiber.io)

## Build Note
- Version 1.2.2
- Go 1.23
- Fiber 2.52.5

## How to
- Clone the repo
- Copy `.env.example` to `.env`
- Generate API Docs using swagger with `make swag`
- Run the application with `make air`
- Build the application with `make build`

## Notes
- Use `merry` to wrap all error


<!-- env GOOS=linux GOARCH=amd64 go build -->