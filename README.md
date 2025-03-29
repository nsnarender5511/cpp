# crules

A tool for managing and synchronizing Cursor rules across multiple projects.

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