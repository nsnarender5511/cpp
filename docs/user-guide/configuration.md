# Configuration

> ⚙️ **Configuring crules for your environment**

## Configuration Methods

crules can be configured using:

1. Environment variables
2. `.env` file in the project root

## Configuration Options

The following configuration options are available:

| Variable | Description | Default |
|----------|-------------|---------|
| APP_NAME | Application name used for system paths | crules |
| RULES_DIR_NAME | Directory name for cursor rules | .cursor/rules |
| REGISTRY_FILE_NAME | Name of the registry file | registry.json |
| DIR_PERMISSION | Directory permission in octal | 0755 |
| FILE_PERMISSION | File permission in octal | 0644 |
| LOG_LEVEL | Logging level | info |
| LOG_FILE_NAME | Name of the log file | crules.log |

## Environment File

Create a `.env` file in the project root with your desired configuration:

```
APP_NAME=crules
RULES_DIR_NAME=.cursor/rules
REGISTRY_FILE_NAME=registry.json
DIR_PERMISSION=0755
FILE_PERMISSION=0644
LOG_LEVEL=info
LOG_FILE_NAME=crules.log
```

A sample configuration file is provided in `.env.example` in the repository.

## Logging Configuration

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

## File Locations

crules follows platform-specific conventions for storing files. The actual paths will respect the configured APP_NAME value:

### Windows
- Configuration: `%APPDATA%\<APP_NAME>`
- Rules & Registry: `%LOCALAPPDATA%\<APP_NAME>\<RULES_DIR_NAME>`
- Logs: `%LOCALAPPDATA%\<APP_NAME>\Logs\<LOG_FILE_NAME>`

### macOS
- Configuration: `~/Library/Application Support/<APP_NAME>`
- Rules & Registry: `~/Library/Application Support/<APP_NAME>/<RULES_DIR_NAME>`
- Logs: `~/Library/Logs/<APP_NAME>/<LOG_FILE_NAME>`

### Linux/Unix
- Configuration: `~/.config/<APP_NAME>`
- Rules & Registry: `~/.local/share/<APP_NAME>/<RULES_DIR_NAME>`
- Logs: `~/.local/state/<APP_NAME>/logs/<LOG_FILE_NAME>`

## Command-Line Options

crules supports the following command-line options that affect the runtime behavior:

- `--verbose`: Show informational messages on console
- `--debug`: Show debug messages on console (implies verbose)

Example:
```bash
crules --debug init
```

## Related Documentation

- [Installation Guide](installation.md)
- [Commands Reference](commands.md)
