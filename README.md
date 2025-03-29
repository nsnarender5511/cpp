# crules

A tool for managing and synchronizing Cursor rules across multiple projects.

## Installation

### macOS

Using Homebrew:

```bash
brew install username/tap/crules
```

### Linux

#### Using pre-built binaries

```bash
# Download and extract
curl -L https://github.com/username/crules/releases/latest/download/crules_Linux_x86_64.tar.gz | tar xz
# Move to a directory in your PATH
sudo mv crules /usr/local/bin/
```

#### Using packages (Debian/Ubuntu)

```bash
# Download the .deb package
curl -LO https://github.com/username/crules/releases/latest/download/crules_amd64.deb
# Install the package
sudo dpkg -i crules_amd64.deb
```

#### Using packages (RHEL/Fedora)

```bash
# Download the .rpm package
curl -LO https://github.com/username/crules/releases/latest/download/crules_amd64.rpm
# Install the package
sudo rpm -i crules_amd64.rpm
```

### Windows

#### Using Scoop

```powershell
# Add the bucket
scoop bucket add username https://github.com/username/scoop-bucket
# Install crules
scoop install crules
```

#### Manual installation

1. Download the latest Windows archive from the [Releases page](https://github.com/username/crules/releases)
2. Extract the archive
3. Add the location to your PATH or move the binary to a directory in your PATH

### Building from source

```bash
# Clone the repository
git clone https://github.com/username/crules.git
cd crules

# Build the binary
go build -o crules cmd/main.go

# Make it executable (macOS/Linux)
chmod +x crules

# Move to a directory in your PATH (optional)
sudo mv crules /usr/local/bin/
```

## Configuration

crules can be configured using environment variables. Create a `.env` file in the project root with any of the following variables:

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name used for system paths | crules |
| RULES_DIR_NAME | Directory name for cursor rules | .cursor/rules |
| REGISTRY_FILE_NAME | Name of the registry file | registry.json |
| DIR_PERMISSION | Directory permission in octal | 0755 |
| FILE_PERMISSION | File permission in octal | 0644 |
| LOG_LEVEL | Logging level | info |
| LOG_FILE_NAME | Name of the log file | crules.log |

Example `.env` file:
```
APP_NAME=crules
RULES_DIR_NAME=.cursor/rules
REGISTRY_FILE_NAME=registry.json
DIR_PERMISSION=0755
FILE_PERMISSION=0644
LOG_LEVEL=info
LOG_FILE_NAME=crules.log
```

A sample configuration file is provided in `.env.example`.

### Logging Levels

The following log levels are available (from most to least verbose):

- `trace`: Extremely detailed information
- `debug`: Detailed information for debugging purposes
- `info`: General information about program execution (default)
- `warn`: Potentially harmful situations that don't prevent operation
- `error`: Error events that might still allow the application to continue
- `fatal`: Severe error events that cause the application to terminate
- `panic`: Severe error events that cause the application to panic

The application uses a dedicated lightweight logger based on Logrus for structured, colorized logging.

### File Locations

crules follows platform-specific conventions for storing files. The actual paths will respect the configured APP_NAME value:

#### Windows
- Configuration: `%APPDATA%\<APP_NAME>`
- Rules & Registry: `%LOCALAPPDATA%\<APP_NAME>\<RULES_DIR_NAME>`
- Logs: `%LOCALAPPDATA%\<APP_NAME>\Logs\<LOG_FILE_NAME>`

#### macOS
- Configuration: `~/Library/Application Support/<APP_NAME>`
- Rules & Registry: `~/Library/Application Support/<APP_NAME>/<RULES_DIR_NAME>`
- Logs: `~/Library/Logs/<APP_NAME>/<LOG_FILE_NAME>`

#### Linux/Unix
- Configuration: `~/.config/<APP_NAME>`
- Rules & Registry: `~/.local/share/<APP_NAME>/<RULES_DIR_NAME>`
- Logs: `~/.local/state/<APP_NAME>/logs/<LOG_FILE_NAME>`

## Usage

```bash
# Show version information
crules --version

# Initialize current directory with rules from main location
crules init

# Merge current rules to main location and sync to all locations
crules merge

# Force sync from main location to current directory
crules sync

# Display all registered projects
crules list

# Remove non-existent projects from registry
crules clean

# Show more detailed output
crules --verbose <command>

# Show debug information
crules --debug <command>
``` 