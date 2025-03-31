# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- New `SourceFolder` option for `cursor++ sync` to allow selecting specific folders from a cloned repository

## [v1.0.0] - 2023-03-29

### 🎉 Major Changes

- **Complete rebranding from "crules" to "cursor++"** to better align with the Cursor IDE integration
- Added GoReleaser for automated releases
- Improved UI/UX for terminal displays
- Enhanced agent selection experience
- Multi-agent mode support for collaborative AI workflows

### ✨ New Features

- Support for homebrew installation on macOS
- Support for Windows and Linux installations
- Integration with `.cursor` directory for seamless IDE experience
- Enhanced documentation with detailed examples
- Interactive agent selection with rich terminal UI

### 🛠️ Improvements

- Simplified command structure
- Improved error handling and user messaging
- Enhanced terminal width detection for better display
- Optimized agent loading for faster startup
- Better Git integration for rule management

### 📚 Documentation

- Complete overhaul of documentation structure
- New examples and tutorials
- Architecture documentation with diagrams
- API documentation for developers
- Enhanced troubleshooting guides

### 🔧 Bug Fixes

- Fixed terminal width detection on some platforms
- Resolved issues with path handling on Windows
- Fixed agent loading in certain edge cases
- Corrected configuration file handling
- Addressed synchronization issues between projects

### 🔄 Coming Soon

- Web-based management interface
- Enhanced multi-agent collaboration features
- Cloud synchronization of rules
- IDE plugin enhancements

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