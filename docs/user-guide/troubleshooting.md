# Troubleshooting Guide

> 🛠️ This guide helps you diagnose and fix common issues with the crules tool.

## Common Issues and Solutions

### Installation Issues

#### Command Not Found

**Problem**: After installation, running `crules` gives "command not found" error.

**Solutions**:

1. **Check installation**:
   ```bash
   which crules  # On macOS/Linux
   where crules  # On Windows
   ```

2. **Add to PATH**: Ensure the installation directory is in your PATH:
   ```bash
   # On macOS/Linux, add to ~/.bashrc or ~/.zshrc
   export PATH=$PATH:/usr/local/bin   # Or wherever crules is installed
   
   # On Windows, add to system PATH through System Properties
   ```

3. **Reinstall**:
   ```bash
   # Using Go
   go install github.com/yourusername/crules/cmd/crules@latest
   
   # Using Homebrew
   brew reinstall yourusername/tap/crules
   ```

#### Permission Denied

**Problem**: Getting "permission denied" when trying to run crules.

**Solutions**:

1. **Make executable** (macOS/Linux):
   ```bash
   chmod +x /path/to/crules
   ```

2. **Run as administrator** (Windows): Right-click and select "Run as administrator".

### Initialization Issues

#### Initialization Fails

**Problem**: The `crules init` command fails with an error.

**Solutions**:

1. **Check permissions**:
   ```bash
   # Ensure you have write permissions to the current directory
   ls -la
   ```

2. **Check configuration**:
   ```bash
   # On macOS/Linux
   cat ~/.config/crules/config.json
   
   # On Windows (PowerShell)
   Get-Content "$env:APPDATA\crules\config.json"
   ```

3. **Reset configuration**:
   ```bash
   # On macOS/Linux
   rm ~/.config/crules/config.json
   
   # On Windows (PowerShell)
   Remove-Item "$env:APPDATA\crules\config.json"
   ```

#### Cannot Find Main Location

**Problem**: The `crules init` command fails with "cannot find main location" error.

**Solutions**:

1. **Set main location**:
   ```bash
   export CRULES_MAIN_LOCATION=/path/to/main/location
   crules init
   ```

2. **Create main location**:
   ```bash
   mkdir -p /path/to/main/location
   export CRULES_MAIN_LOCATION=/path/to/main/location
   crules init
   ```

### Synchronization Issues

#### Sync Fails

**Problem**: The `crules sync` command fails to synchronize rules.

**Solutions**:

1. **Check permissions**:
   ```bash
   # Ensure you have read permissions for the main location
   ls -la $CRULES_MAIN_LOCATION
   ```

2. **Verify configuration**:
   ```bash
   crules list  # Check if current project is registered
   ```

3. **Force initialization**:
   ```bash
   crules init --force
   ```

#### Files Not Syncing

**Problem**: Files aren't being synced between locations.

**Solutions**:

1. **Check rules directory**:
   ```bash
   # Verify the rules directory exists in both locations
   ls -la .cursor/rules
   ls -la $CRULES_MAIN_LOCATION/.cursor/rules
   ```

2. **Force sync**:
   ```bash
   crules sync --force
   ```

3. **Check for conflicts**:
   ```bash
   # Look for .conflict files
   find .cursor/rules -name "*.conflict"
   ```

### Agent Issues

#### Agent List Empty

**Problem**: The `crules agent list` command shows no agents.

**Solutions**:

1. **Check rules directory**:
   ```bash
   # Verify agent definition files exist
   ls -la .cursor/rules/*.mdc
   ```

2. **Add sample agents**:
   ```bash
   # Copy sample agent definitions to your rules directory
   cp /path/to/sample/agents/*.mdc .cursor/rules/
   ```

3. **Sync from main location**:
   ```bash
   crules sync
   ```

#### Agent Not Found When Using Numeric Index

**Problem**: The `crules agent info <number>` command returns "Agent not found" even though the number is valid in the list.

**Solutions**:

1. **Rebuild the application**:
   ```bash
   # Ensure you have the latest version of the binary
   make build
   ```

2. **Verify agent listing**:
   ```bash
   # Confirm the number corresponds to a valid agent
   crules agent
   ```

3. **Use string ID instead**:
   ```bash
   # Reference the agent by its string ID rather than position number
   crules agent info <agent-id>
   ```

#### Agent Selection Fails

**Problem**: The `crules agent select` command fails.

**Solutions**:

1. **Check terminal capabilities**:
   ```bash
   # Some terminal emulators may not support the interactive UI
   crules agent select --simple  # Use simple selection mode
   ```

2. **Directly select an agent**:
   ```bash
   # Set the selected agent directly
   crules agent info <agent-id>  # First find an available agent ID
   export CRULES_SELECTED_AGENT=<agent-id>
   ```

### Configuration Issues

#### Configuration File Corrupted

**Problem**: The configuration file is corrupted or has invalid JSON.

**Solutions**:

1. **Check file contents**:
   ```bash
   cat ~/.config/crules/config.json
   ```

2. **Reset configuration**:
   ```bash
   rm ~/.config/crules/config.json
   crules init
   ```

3. **Restore from backup**:
   ```bash
   cp ~/crules-config-backup.json ~/.config/crules/config.json
   ```

#### Cannot Access Configuration

**Problem**: The tool cannot access or modify the configuration file.

**Solutions**:

1. **Check permissions**:
   ```bash
   ls -la ~/.config/crules/
   ```

2. **Fix permissions**:
   ```bash
   chmod 755 ~/.config/crules/
   chmod 644 ~/.config/crules/config.json
   ```

3. **Create directory**:
   ```bash
   mkdir -p ~/.config/crules/
   ```

## Diagnostic Tools

### Version Check

```bash
crules --version
```

This shows the version information, which is useful when reporting issues.

### Verbose Mode

```bash
crules --verbose <command>
```

This runs a command with more detailed output, which can help diagnose issues.

### Debug Mode

```bash
crules --debug <command>
```

This runs a command with debug-level logging, which provides even more information for troubleshooting.

### Log Files

Check the log files for detailed error information:

```bash
# On macOS
cat ~/Library/Logs/crules/crules.log

# On Linux
cat ~/.local/state/crules/logs/crules.log

# On Windows (PowerShell)
Get-Content "$env:LOCALAPPDATA\crules\Logs\crules.log"
```

## Common Error Codes

| Exit Code | Description | Possible Solution |
|-----------|-------------|-------------------|
| 1 | General error | Check command syntax and try again |
| 2 | Configuration error | Check or reset configuration file |
| 3 | File system error | Check permissions and file existence |
| 4 | Network error | Check your network connection |
| 5 | User input error | Check your command arguments |

## Reporting Issues

If you've tried the solutions above and are still experiencing issues, you can:

1. **Check existing issues** on the [GitHub repository](https://github.com/yourusername/crules/issues)
2. **Create a new issue** with the following information:
   - crules version (`crules --version`)
   - Operating system and version
   - Command you were trying to run
   - Error message or unexpected behavior
   - Steps to reproduce the issue
   - Any relevant logs or screenshots

## Contact Support

For additional support, you can:

- Join the [Discord community](https://discord.gg/yourusername)
- Email support at support@yourusername.com
- Open a discussion on the [GitHub repository](https://github.com/yourusername/crules/discussions)

## Next Steps

If you've resolved your issue, you can continue with:

- [Installation Guide](./installation.md)
- [Configuration Guide](./configuration.md)
- [Agent System](./agents.md)
- [Command Reference](./commands.md)
