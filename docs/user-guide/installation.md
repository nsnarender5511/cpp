# Installation Guide

> 📥 **Instructions for installing crules on different platforms**

## Prerequisites

Before installing crules, ensure you have:

- Go 1.16 or later (for building from source)
- Git (for downloading the source)

## Installing from Binary

### macOS

```bash
# Download the latest release
curl -L https://github.com/yourusername/crules/releases/latest/download/crules-darwin-amd64 -o crules

# Make it executable
chmod +x crules

# Move to a directory in your PATH
sudo mv crules /usr/local/bin/
```

### Linux

```bash
# Download the latest release
curl -L https://github.com/yourusername/crules/releases/latest/download/crules-linux-amd64 -o crules

# Make it executable
chmod +x crules

# Move to a directory in your PATH
sudo mv crules /usr/local/bin/
```

### Windows

1. Download the latest release from the [Releases page](https://github.com/yourusername/crules/releases)
2. Rename the file to `crules.exe`
3. Move the file to a directory in your PATH, or add its location to your PATH

## Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/crules.git
   cd crules
   ```

2. Build the binary:
   ```bash
   go build -o crules cmd/main.go
   ```

3. Make it executable (macOS/Linux):
   ```bash
   chmod +x crules
   ```

4. Move to a directory in your PATH (optional):
   ```bash
   sudo mv crules /usr/local/bin/
   ```

## Verifying the Installation

To verify that crules is installed correctly, run:

```bash
crules --version
```

You should see the version information displayed.

## Next Steps

After installation:

1. [Configure crules](configuration.md) with your preferred settings
2. Learn about available [commands](commands.md)
3. Check out the [basic usage examples](../examples/basic-usage.md)
