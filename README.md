# Puasa Sunnah / Sunnah Fasting API

<div align="center">

[![Better Uptime Badge](https://betteruptime.com/status-badges/v1/monitor/ds3l.svg)](https://betteruptime.com/?utm_source=status_badge)
![Deployment](https://img.shields.io/github/workflow/status/granitebps/puasa-sunnah-api/deployment/main)

</div>

## Link
- [API](https://api.puasa-sunnah.granitebps.com)
- [API Docs](https://api.puasa-sunnah.granitebps.com/swagger)

## Framework
- [Fiber](https://gofiber.io)

## Build Note
- Version 0.0.2
- Go 1.17.3
- Fiber 2.23.0

## How to
- Clone the repo
- Copy `.env.example` to `.env`
- Generate API Docs using swagger with `make doc`
- Run the application with `make start` or with `make air` to use hot reload
- Build the application with `make build`