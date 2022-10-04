# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.2] - 2022-10-04

### Fixed
- #15 Fixed documentation
- #20 Update go-keyhub dependency to fix max 100 items bug

## [1.2.1] - 2022-09-21

### Added
- Add build for Linux-Arm64

## [1.2.0] - 2022-09-13

### Added
- add nested_under_groupuuid parameter

### Changed
- Update documentation
- Members are optional if a nested_under_groupuuid has been provided
- Version bump dependencies go-keyhub to 1.2.1

## [1.1.0] - 2022-07-22

### Added
- Group resource
- Group datasource
- Provisionedsystem resource
- GroupOnSystem resource
- ClientApplication resource