# Installation Guide

> 📥 This guide explains how to install crules on different platforms.

## Overview

crules is available for macOS, Linux, and Windows. There are several installation methods depending on your platform and preferences.

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

The fastest way to install crules is using the install script:

```bash
curl -sL https://github.com/org/crules/releases/latest/download/install.sh | bash
```

This script will:
1. Detect your operating system
2. Download the appropriate binary
3. Install it to `~/.local/bin` (or `/usr/local/bin` if run with sudo)
4. Add it to your PATH (if needed)

<details>
  <summary>📺 View Installation Process</summary>
  <img src="../assets/gifs/installation/install-script.gif" alt="Install script process" width="600">
</details>

### Method 2: Homebrew (macOS)

If you use Homebrew on macOS, you can install crules with:

```bash
brew install crules
```

### Method 3: Manual Installation

You can manually download and install the binary for your platform:

#### macOS

```bash
# Download the latest macOS binary
curl -Lo crules https://github.com/org/crules/releases/latest/download/crules-darwin-amd64

# Make it executable
chmod +x crules

# Move it to a directory in your PATH
sudo mv crules /usr/local/bin/
```

#### Linux

```bash
# Download the latest Linux binary
curl -Lo crules https://github.com/org/crules/releases/latest/download/crules-linux-amd64

# Make it executable
chmod +x crules

# Move it to a directory in your PATH
sudo mv crules /usr/local/bin/
```

#### Windows

1. Download the [latest Windows binary](https://github.com/org/crules/releases/latest/download/crules-windows-amd64.exe)
2. Rename it to `crules.exe`
3. Move it to a directory in your PATH (e.g., `C:\Windows\System32\`)

Alternatively, using PowerShell:

```powershell
# Download the latest Windows binary
Invoke-WebRequest -Uri https://github.com/org/crules/releases/latest/download/crules-windows-amd64.exe -OutFile crules.exe

# Move it to a directory in your PATH
Move-Item -Path crules.exe -Destination "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\"
```

### Method 4: Building from Source

If you prefer to build from source, follow these steps:

1. Ensure you have Go 1.19+ installed
2. Clone the repository:
   ```bash
   git clone https://github.com/org/crules.git
   ```
3. Navigate to the directory:
   ```bash
   cd crules
   ```
4. Build the binary:
   ```bash
   go build -o crules ./cmd/crules
   ```
5. Install it:
   ```bash
   sudo mv crules /usr/local/bin/
   ```

For detailed build instructions, see the [Building from Source](../developer-guide/building.md) guide.

## Verifying Installation

After installation, verify that crules is installed correctly:

```bash
crules --version
```

You should see output similar to:

```
crules v0.1.0
```

## First-Time Setup

After installing crules, the first time you run it, you'll need to set up your main rules location. This is where all your rules will be stored and synced from.

```bash
crules init
```

The first time you run this command, you'll be guided through the setup process. See the [Getting Started Guide](./getting-started.md) for more information.

## Updating

### Using the Install Script

To update crules using the install script:

```bash
curl -sL https://github.com/org/crules/releases/latest/download/install.sh | bash
```

### Using Homebrew

If you installed using Homebrew:

```bash
brew upgrade crules
```

### Manual Update

To update manually, download the latest binary for your platform and replace the existing binary.

## Uninstalling

### macOS and Linux

```bash
# Remove the binary
sudo rm /usr/local/bin/crules

# Remove configuration and data (optional)
rm -rf ~/.config/crules
rm -rf ~/.local/share/crules
```

### Windows

```powershell
# Remove the binary
Remove-Item "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps\crules.exe"

# Remove configuration and data (optional)
Remove-Item -Recurse -Force "$env:APPDATA\crules"
```

## Troubleshooting Installation

If you encounter any issues during installation, please refer to the [Troubleshooting Guide](./troubleshooting.md#installation-issues) for common problems and solutions.

You can also open an issue on our [GitHub repository](https://github.com/nsnarender5511/crules/issues) for additional help.

## See Also

- [Getting Started Guide](./getting-started.md): Quick introduction to crules
- [Command Reference](./commands.md): Detailed information about all commands
- [Configuration Guide](./configuration.md): How to configure crules

## Navigation

- Next: [Getting Started](./getting-started.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
