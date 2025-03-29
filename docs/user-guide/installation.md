# Installation Guide

> 📥 This guide explains how to install and set up the crules tool on different operating systems.

## System Requirements

Before installing crules, ensure your system meets the following requirements:

- **Operating System**: Windows, macOS, or Linux
- **Disk Space**: At least 20MB of free disk space
- **Dependencies**: Git (for certain operations)

## Installation Methods

### Using Go

If you have Go 1.20 or later installed, you can install crules directly:

```bash
go install github.com/yourusername/crules/cmd/crules@latest
```

This command will download the source code, compile it, and install the binary in your `$GOPATH/bin` directory. Make sure this directory is in your system's PATH.

### From Binary Releases

#### macOS

1. **Using Homebrew**:

```bash
# Add the tap
brew tap yourusername/tap
# Install crules
brew install yourusername/tap/crules
```

2. **Manual Installation**:

```bash
# Download the latest macOS binary
curl -L https://github.com/yourusername/crules/releases/latest/download/crules_darwin_amd64.tar.gz -o crules.tar.gz

# Extract the archive
tar -xzf crules.tar.gz

# Move to a directory in your PATH
sudo mv crules /usr/local/bin/

# Make it executable
chmod +x /usr/local/bin/crules

# Clean up
rm crules.tar.gz
```

#### Linux

1. **Using pre-built binaries**:

```bash
# Download and extract
curl -L https://github.com/yourusername/crules/releases/latest/download/crules_linux_amd64.tar.gz | tar xz

# Move to a directory in your PATH
sudo mv crules /usr/local/bin/

# Make it executable
chmod +x /usr/local/bin/crules
```

2. **Using packages (Debian/Ubuntu)**:

```bash
# Download the .deb package
curl -LO https://github.com/yourusername/crules/releases/latest/download/crules_amd64.deb

# Install the package
sudo dpkg -i crules_amd64.deb
```

3. **Using packages (RHEL/Fedora)**:

```bash
# Download the .rpm package
curl -LO https://github.com/yourusername/crules/releases/latest/download/crules_amd64.rpm

# Install the package
sudo rpm -i crules_amd64.rpm
```

#### Windows

1. **Using Scoop**:

```powershell
# Add the bucket
scoop bucket add yourusername https://github.com/yourusername/scoop-bucket

# Install crules
scoop install crules
```

2. **Manual installation**:

- Download the latest Windows archive from the [Releases page](https://github.com/yourusername/crules/releases)
- Extract the archive
- Add the location to your PATH or move the binary to a directory in your PATH

### Building from source

If you prefer to build from source:

```bash
# Clone the repository
git clone https://github.com/yourusername/crules.git
cd crules

# Build the binary
go build -o crules ./cmd/crules

# Make it executable (macOS/Linux)
chmod +x crules

# Move to a directory in your PATH (optional)
sudo mv crules /usr/local/bin/
```

## Verifying the Installation

To verify that crules is installed correctly:

```bash
crules --version
```

This should display the version information of the installed crules tool.

## First-time Setup

After installing crules, you should initialize it in your project:

```bash
# Navigate to your project directory
cd your-project

# Initialize crules
crules init
```

This will set up the necessary directory structure and configuration files for crules.

## Updating

### Using Go

If you installed using Go, update with:

```bash
go install github.com/yourusername/crules/cmd/crules@latest
```

### Using package managers

If you installed using a package manager, update through that package manager:

```bash
# Homebrew
brew upgrade yourusername/tap/crules

# Scoop
scoop update crules
```

### From binary releases

For manual installations, download the latest release and replace your existing binary.

## Troubleshooting

If you encounter issues during installation, see the [Troubleshooting Guide](./troubleshooting.md) for common problems and solutions.

## Next Steps

Now that you have installed crules, you can:

- Learn about [Configuration](./configuration.md)
- Explore the [Agent System](./agents.md)
- Review the [Command Reference](./commands.md)
- Try out some [Examples](../examples/agent-workflows.md)
