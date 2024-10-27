# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

You can find and compare releases at the GitHub release page.

## [Unreleased]

## [1.2.2] - 2024-10-27

### Added
- Add CI/CD for auto deployment
- Add background color and text color to type

## [1.2.1] - 2024-10-24

### Added
- Add request timeout
- Add response cache using redis
- Add delete API to fasting, category, type and source

### Changed
- Update default filter for fasting

## [1.2.0] - 2024-05-01

### Fixed
- Order fasting by date
- Add check between day,month,year to date

### Added
- Add database backup
- Add human date to fasting api
- Add 2023 data
- Add 2024 data
- Add test

## [1.1.0] - 2024-03-07

### Changed
- Update to Golang version 1.21
- Another refactoring

### Added
- Add Gorm to connect to MySQL

## [1.0.0] - 2023-11-16

### Changed
- Massive refactoring

## [0.1.0] - 2022-05-03

### Added
- Add category and type data to fasting response API
- Add filter to fasting response API
- Add sentry
- Add testing
- Add CI

### Changed
- Improve security

## [0.0.2] - 2022-04-20

### Added
- Add Makefile file
- Add README file
- Add github action for deployment
- Add CHANGELOG file
- Add Swagger docs
- Add log file
- Add hot reload

## [0.0.1] - 2022-04-18

### Added
- Create JSON file for all data
- Setup fiber as web framework
- Setup sources API
- Setup categories API
- Setup types API
- Setup fasting/puasa-sunnah API
- Setup env file
- Prepare for deployment

[Unreleased]: https://github.com/granitebps/puasa-sunnah-api/compare/main...dev
[1.2.2]: https://github.com/granitebps/puasa-sunnah-api/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/granitebps/puasa-sunnah-api/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/granitebps/puasa-sunnah-api/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/granitebps/puasa-sunnah-api/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/granitebps/puasa-sunnah-api/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/granitebps/puasa-sunnah-api/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/granitebps/puasa-sunnah-api/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/granitebps/puasa-sunnah-api/releases/tag/v0.0.1