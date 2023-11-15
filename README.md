# Puasa Sunnah / Sunnah Fasting API

<div align="center">

[![Better Uptime Badge](https://betteruptime.com/status-badges/v1/monitor/ds3l.svg)](https://betteruptime.com/?utm_source=status_badge)
[![CircleCI](https://circleci.com/gh/granitebps/puasa-sunnah-api/tree/main.svg?style=shield)](https://circleci.com/gh/granitebps/puasa-sunnah-api/tree/main)

</div>

## Link
- [API](https://api.puasa-sunnah.granitebps.com)
- [API Docs](https://api.puasa-sunnah.granitebps.com/swagger)

## Framework
- [Fiber](https://gofiber.io)

## Build Note
- Version 0.1.0
- Go 1.17.3
- Fiber 2.32.0

## How to
- Clone the repo
- Copy `.env.example` to `.env`
- Generate API Docs using swagger with `swag init`
- Run the application with `make start` or with `air` to use hot reload
- Build the application with `make build`

## Notes
- Use `merry` to wrap all error


<!-- env GOOS=linux GOARCH=amd64 go build -->

### TODO
- [] Refactor
- [] Update response helper and log if 500
- [x] Log file
- [x] helper error using merry and fiber.NewError