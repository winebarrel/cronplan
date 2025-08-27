# Changelog

## [2.0.2] - 2025-08-28

### Fixed

* Fix .goreleaser.yml for notarizing. [pull#111](https://github.com/winebarrel/cronplan/pull/111)

## [2.0.1] - 2025-08-27

### Fixed

* Fix package. Add "v2". [pull#109](https://github.com/winebarrel/cronplan/pull/109)

## [2.0.0] - 2025-08-27

### Fixed

* **Breaking change**: Fix "31W" behavior. [pull#107](https://github.com/winebarrel/cronplan/pull/107)

## [1.10.5] - 2025-01-26

### Fixed

* Fix error message. [pull#87](https://github.com/winebarrel/cronplan/pull/87)

## [1.10.4] - 2024-11-11

### Fixed

* Fix cronskd `-e` option.


## [1.10.3] - 2024-11-10

### Added

* Add cronskd command.

## [1.10.2] - 2024-11-09

### Changed

* Remove unnecessary modules from go.mod.

## [1.10.1] - 2023-10-01

### Added

* Add crongrep CLI.

## [1.10.0] - 2023-10-01

### Added

* Support "LW" in day-of-month. (e.g "0 0 LW * ? *")

## [1.9.2] - 2023-09-30

### Fixed

* Fix bug for "L" without wday in day-of-week.

## [1.9.1] - 2023-09-29

### Added

* Support "L" without wday in day-of-week. (e.g "0 0 ? * L *")

## [1.9.0] - 2023-09-29

### Changed

* Enable govet.

### Fixed

* Fix for last weekday of the month. (e.g. "6L", "MONL")

## [1.8.1] - 2023-06-24

### Added

* Add Changelog.

### Changed

* Enable gofmt,misspell.
* Merge goreleaser configs.

## [1.8.0] - 2023-06-20

### Added

* Add license.

<!-- cf. https://keepachangelog.com/ -->
