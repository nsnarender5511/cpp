# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-03-29

### Added
- Added GoReleaser for automated multi-platform releases
- Added support for Homebrew installation
- Added support for Scoop installation
- Added support for Linux package managers (deb/rpm)
- Added version flag to CLI
- Added gitignore management functionality

### Changed
- Updated documentation with installation instructions for all platforms

### Initial Features
- Core synchronization functionality
- Registry management for multiple projects
- Platform-specific path handling
- Configuration via environment variables
- Logging system
- Commands: init, merge, sync, list, clean 