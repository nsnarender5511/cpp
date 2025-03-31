# Installation Guide

> 📥 This guide explains how to install cursor++ on different platforms.

## Overview

cursor++ is available for macOS, Linux, and Windows. There are several installation methods depending on your platform and preferences.

## System Requirements

- **Operating System**:
  - macOS 10.15 (Catalina) or later
  - Linux (Ubuntu 18.04+, Debian 10+, Fedora 32+, or compatible distributions)
  - Windows 10 or later

- **Disk Space**: Minimum 10MB free space
- **Optional Dependencies**:
  - Git (for repository-based rule management)

## Installation Methods

### Method 1: Using the Install Script (macOS and Linux)

The fastest way to install cursor++ is using the install script:

```bash
curl -sL https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/install.sh | bash
```

This script will:
1. Detect your operating system
2. Download the appropriate binary
3. Install it to `~/.local/bin` (or `/usr/local/bin` if run with sudo)
4. Add it to your PATH (if needed)

### Method 2: Homebrew (macOS)

If you use Homebrew on macOS, you can install cursor++ with:

```bash
brew install cursor++
```

### Method 3: Manual Installation

You can manually download and install the binary for your platform:

#### macOS

```bash
# Download the latest macOS binary
curl -Lo cursor++ https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/cursor++-darwin-amd64

# Make it executable
chmod +x cursor++

# Move it to a directory in your PATH
sudo mv cursor++ /usr/local/bin/
```

#### Linux

```bash
# Download the latest Linux binary
curl -Lo cursor++ https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/cursor++-linux-amd64

# Make it executable
chmod +x cursor++

# Move it to a directory in your PATH
sudo mv cursor++ /usr/local/bin/
```

#### Windows

1. Download the [latest Windows binary](https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/cursor++-windows-amd64.exe)
2. Rename it to `cursor++.exe`
3. Move it to a directory in your PATH (e.g., `C:\Windows\System32\`)

Alternatively, using PowerShell:

```powershell
# Download the latest Windows binary
Invoke-WebRequest -Uri https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/cursor++-windows-amd64.exe -OutFile cursor++.exe

# Move it to a directory in your PATH
Move-Item -Path cursor++.exe -Destination "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\"
```

### Method 4: Building from Source

If you prefer to build from source, follow these steps:

1. Ensure you have Go 1.23+ installed
2. Clone the repository:
   ```bash
   git clone https://github.com/cursor-ai/cursor-plus-plus.git
   ```
3. Navigate to the directory:
   ```bash
   cd cursor-plus-plus
   ```
4. Build the binary:
   ```bash
   go build -o cursor++ ./cmd
   ```
5. Install it:
   ```bash
   sudo mv cursor++ /usr/local/bin/
   ```

For detailed build instructions, see the [Building from Source](../developer-guide/building.md) guide.

## Verifying Installation

After installation, verify that cursor++ is installed correctly:

```bash
cursor++ --version
```

You should see output similar to:

```
cursor++ v1.0.0
```

## First-Time Setup

After installing cursor++, the first time you run it, you'll need to set up your main rules location. This is where all your rules will be stored and synced from.

```bash
cursor++ init
```

The first time you run this command, you'll be guided through the setup process. See the [Getting Started Guide](./getting-started.md) for more information.

## Updating

### Using the Install Script

To update cursor++ using the install script:

```bash
curl -sL https://github.com/cursor-ai/cursor-plus-plus/releases/latest/download/install.sh | bash
```

### Using Homebrew

If you installed using Homebrew:

```bash
brew upgrade cursor++
```

### Manual Update

To update manually, download the latest binary for your platform and replace the existing binary.

## Uninstalling

### macOS and Linux

```bash
# Remove the binary
sudo rm /usr/local/bin/cursor++

# Remove configuration and data (optional)
rm -rf ~/.config/cursor++
rm -rf ~/.local/share/cursor++
```

### Windows

```powershell
# Remove the binary
Remove-Item "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\cursor++.exe"

# Remove configuration and data (optional)
Remove-Item -Recurse -Force "$env:APPDATA\cursor++"
```

## Troubleshooting Installation

If you encounter any issues during installation, please refer to the [Troubleshooting Guide](./troubleshooting.md#installation-issues) for common problems and solutions.

You can also open an issue on our [GitHub repository](https://github.com/cursor-ai/cursor-plus-plus/issues) for additional help.

## See Also

- [Getting Started Guide](./getting-started.md): Quick introduction to cursor++
- [Command Reference](./commands.md): Detailed information about all commands
- [Configuration Guide](./configuration.md): How to configure cursor++

## Navigation

- Next: [Getting Started](./getting-started.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
