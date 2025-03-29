# Troubleshooting

> ❓ **Solutions for common issues with crules**

## Common Issues and Solutions

### Command Not Found

**Issue**: `crules: command not found`

**Possible Solutions**:
1. Ensure that crules is installed correctly:
   ```bash
   which crules
   ```
2. If it's not in your PATH, reinstall it by following the [installation guide](installation.md).
3. If you built from source, make sure you moved the binary to a directory in your PATH.

### Permission Denied

**Issue**: `Permission denied` when running crules commands

**Possible Solutions**:
1. Ensure the binary has execute permissions:
   ```bash
   chmod +x /path/to/crules
   ```
2. Check if you have write permissions to the target directories:
   ```bash
   # For the current project
   ls -la .cursor/rules
   
   # For the main location (based on your OS)
   # macOS example:
   ls -la ~/Library/Application\ Support/crules
   ```

### Rules Not Syncing

**Issue**: Changes to rules are not being synced across projects

**Possible Solutions**:
1. Make sure all projects are properly registered:
   ```bash
   crules list
   ```
2. Verify that the rules directory exists in your project:
   ```bash
   ls -la .cursor/rules
   ```
3. Try running with debug output to see what's happening:
   ```bash
   crules --debug sync
   ```
4. Check the logs for more detailed information (see [logging configuration](configuration.md)).

### Registry Errors

**Issue**: Errors related to registry operations

**Possible Solutions**:
1. Check if the registry file exists and is valid:
   ```bash
   # Path depends on your OS, see the configuration guide
   cat ~/Library/Application\ Support/crules/.cursor/rules/registry.json
   ```
2. If the registry is corrupted, try cleaning it:
   ```bash
   crules clean
   ```
3. In extreme cases, you might need to manually remove the registry file and reinitialize your projects.

### Path Issues in Windows

**Issue**: Problems with file paths on Windows systems

**Possible Solutions**:
1. Ensure paths don't contain special characters
2. Try using forward slashes (`/`) instead of backslashes (`\`) in configuration
3. Make sure long path support is enabled in Windows

## Log Analysis

When troubleshooting, check the application logs:

- **Windows**: `%LOCALAPPDATA%\crules\Logs\crules.log`
- **macOS**: `~/Library/Logs/crules/crules.log`
- **Linux**: `~/.local/state/crules/logs/crules.log`

Increase logging detail by setting the LOG_LEVEL in your .env file:
```
LOG_LEVEL=debug
```

## Debugging Tips

1. **Use debug mode** for more verbose output:
   ```bash
   crules --debug <command>
   ```

2. **Check environment variables** to ensure proper configuration:
   ```bash
   # Linux/macOS
   env | grep -i crules
   
   # Windows
   set | findstr /i crules
   ```

3. **Verify file permissions** for both the crules binary and its data directories

## Reporting Issues

If you can't resolve an issue, report it by:

1. Collecting relevant logs
2. Noting exact command and error message
3. Describing your environment (OS, version)
4. Opening an issue in the project repository with this information

## Related Documentation

- [Configuration Guide](configuration.md)
- [Commands Reference](commands.md)
