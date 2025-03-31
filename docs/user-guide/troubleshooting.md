# Troubleshooting Guide

> 🔧 This guide helps you diagnose and resolve common issues that may occur when using crules.

## Common Issues and Solutions

### Installation Issues

#### Tool Not Found After Installation

**Symptoms:**
- Command `crules` not found error
- "Command not found" when trying to run crules

**Solutions:**
1. Check if the binary is in your PATH:
   ```bash
   which crules
   ```

2. Verify installation directory is in your PATH:
   ```bash
   echo $PATH
   ```

3. Reinstall crules:
   ```bash
   curl -sL https://github.com/org/crules/releases/latest/download/install.sh | bash
   ```

4. Add installation directory to your PATH:
   ```bash
   # For Bash (add to ~/.bashrc or ~/.bash_profile)
   export PATH=$PATH:/path/to/crules/bin
   
   # For Zsh (add to ~/.zshrc)
   export PATH=$PATH:/path/to/crules/bin
   ```

#### Permission Denied Error

**Symptoms:**
- "Permission denied" when running crules
- "Cannot execute binary file"

**Solutions:**
1. Make the binary executable:
   ```bash
   chmod +x /path/to/crules
   ```

2. Check file ownership:
   ```bash
   ls -la /path/to/crules
   ```

3. Run with sudo (if necessary):
   ```bash
   sudo crules --version
   ```

### Initialization Issues

#### Main Rules Location Not Set

**Symptoms:**
- Error about main rules location not existing
- "Main rules location does not exist" warning

**Solutions:**
1. Follow the first-time setup prompts to create or clone a main location
2. Set main location manually in config:
   ```json
   // ~/.config/crules/config.json
   {
     "mainLocation": "/path/to/main/rules"
   }
   ```

3. Use environment variable:
   ```bash
   export APP_NAME=custom-name
   crules init
   ```

#### Failed to Create Rules Directory

**Symptoms:**
- "Failed to create directory" error
- Permission-related errors on initialization

**Solutions:**
1. Check directory permissions:
   ```bash
   ls -la ~/.cursor/
   ```

2. Create the directory manually:
   ```bash
   mkdir -p ~/.cursor/rules
   chmod 755 ~/.cursor/rules
   ```

3. Check for disk space issues:
   ```bash
   df -h
   ```

4. Try running with elevated privileges:
   ```bash
   sudo crules init
   ```

### Synchronization Issues

#### Sync Command Fails

**Symptoms:**
- "Failed to sync rules" error
- Files not being updated during sync

**Solutions:**
1. Verify the main location exists and has rules:
   ```bash
   ls -la ~/.cursor/rules
   ```

2. Check file permissions in both locations:
   ```bash
   ls -la ~/.cursor/rules
   ls -la ./.cursor/rules
   ```

3. Run with debug output to see detailed errors:
   ```bash
   crules --debug sync
   ```

4. Try resetting by reinitializing:
   ```bash
   crules init
   ```

#### Merge Command Fails

**Symptoms:**
- "Failed to merge rules" error
- Changes not propagating to main location

**Solutions:**
1. Check if the main location is writable:
   ```bash
   touch ~/.cursor/rules/test && rm ~/.cursor/rules/test
   ```

2. Verify your project is registered:
   ```bash
   crules list
   ```

3. Check for file conflicts:
   ```bash
   crules --debug merge
   ```

4. Try with `--verbose` flag:
   ```bash
   crules --verbose merge
   ```

### Registry Issues

#### Projects Missing from Registry

**Symptoms:**
- Projects not showing in `crules list`
- "Project not registered" errors

**Solutions:**
1. Re-register the project:
   ```bash
   crules init
   ```

2. Check registry file for corruption:
   ```bash
   cat ~/.local/share/crules/registry.json
   ```

3. Reset registry file:
   ```bash
   rm ~/.local/share/crules/registry.json
   ```

4. Run registry repair:
   ```bash
   crules clean
   ```

#### Cannot Clean Non-existent Projects

**Symptoms:**
- Error when running `crules clean`
- Non-existent projects still showing in list

**Solutions:**
1. Check registry file permissions:
   ```bash
   ls -la ~/.local/share/crules/registry.json
   ```

2. Manually edit registry file:
   ```bash
   # Edit registry file directly
   nano ~/.local/share/crules/registry.json
   ```

3. Reset registry completely:
   ```bash
   echo '{"projects":[]}' > ~/.local/share/crules/registry.json
   ```

### Agent System Issues

#### Agent Not Found

**Symptoms:**
- "Agent not found" error
- "No agents available" message

**Solutions:**
1. Check if any agent files exist in any of these locations:
   ```bash
   ls -la ~/.cursor/rules/*.mdc
   ls -la ./.cursor/rules/*.mdc
   ls -la /usr/local/share/cursor-rules/*.mdc
   ```

   > **Note**: Since the v1.1 update, the system automatically checks multiple locations for agent rules: the project directory, home directory, and system-wide directory. This makes it more robust when finding agents.

2. Create a basic agent file:
   ```bash
   echo "# Test Agent\n\nA simple test agent" > ~/.cursor/rules/test-agent.mdc
   ```

3. Sync rules from main location:
   ```bash
   crules sync
   ```

4. Import agents from URL:
   ```bash
   crules import https://cursor.directory/example-agent
   ```

#### Directory Structures Not Being Found

**Symptoms:**
- System doesn't recognize the rules directory
- Path-related errors when accessing rules

**Solutions:**
1. Verify the directory structure exists:
   ```bash
   mkdir -p ~/.cursor/rules
   mkdir -p ./.cursor/rules
   ```

2. Run init with debug mode to see directory detection process:
   ```bash
   crules --debug init
   ```

3. Check permissions on all potential rules directories:
   ```bash
   ls -la ~/.cursor/
   ls -la ./.cursor/
   ls -la /usr/local/share/
   ```

#### Source Folder Not Found

**Symptoms:**
- "Source folder does not exist" error during init
- Failed repository cloning with folder-related errors

**Solutions:**
1. Verify the source folder exists in the repository:
   ```bash
   # Clone the repository temporarily to check its structure
   git clone https://github.com/repository-url.git temp-check
   ls -la temp-check/
   ```

2. Configure a different source folder if needed:
   ```bash
   # Edit config to change source folder name
   echo "SOURCE_FOLDER=new-folder-name" >> ~/.config/cursor++/config.env
   ```
   
3. Check your configuration to see what source folder is currently set:
   ```bash
   grep SOURCE_FOLDER ~/.config/cursor++/config.env
   ```
   
4. Retry with verbose output to see detailed clone and copy steps:
   ```bash
   crules --verbose init
   ```

5. If the repository structure has changed, you may need to manually reorganize:
   ```bash
   # Clone repository to a temporary location
   git clone https://github.com/repository-url.git temp-repo
   # Create the source folder structure if it doesn't exist
   mkdir -p temp-repo/default
   # Move needed files to the source folder
   mv temp-repo/*.mdc temp-repo/default/
   # Commit and push changes if you own the repository
   ```

#### Agent Selection Fails

**Symptoms:**
- Error when running `crules agent select`
- Selection UI crashes or exits unexpectedly

**Solutions:**
1. Try selecting by ID directly:
   ```bash
   crules agent info wizard
   ```

2. Check terminal compatibility:
   ```bash
   echo $TERM
   ```

3. Update terminal settings:
   ```bash
   export TERM=xterm-256color
   ```

4. Run in debug mode:
   ```bash
   crules --debug agent select
   ```

### Configuration Issues

#### Configuration File Corruption

**Symptoms:**
- "Failed to parse config" error
- Unexpected behavior despite correct commands

**Solutions:**
1. Check configuration file:
   ```bash
   cat ~/.config/crules/config.json
   ```

2. Reset configuration file:
   ```bash
   rm ~/.config/crules/config.json
   crules --version  # Creates new default config
   ```

3. Create minimal configuration:
   ```bash
   echo '{"mainLocation":"~/.cursor/rules"}' > ~/.config/crules/config.json
   ```

#### Wrong Permissions on Config Directories

**Symptoms:**
- Unable to write to config or data directories
- Permission denied errors

**Solutions:**
1. Check directory permissions:
   ```bash
   ls -la ~/.config/crules/
   ls -la ~/.local/share/crules/
   ```

2. Fix permissions:
   ```bash
   chmod 755 ~/.config/crules/
   chmod 644 ~/.config/crules/config.json
   chmod 755 ~/.local/share/crules/
   chmod 644 ~/.local/share/crules/registry.json
   ```

3. Check ownership:
   ```bash
   chown -R $(whoami) ~/.config/crules/
   chown -R $(whoami) ~/.local/share/crules/
   ```

## Diagnostic Commands

### Version Check

Check that crules is properly installed and get version information:

```bash
crules --version
```

### Debug Output

Run any command with `--debug` to get detailed diagnostic information:

```bash
crules --debug init
crules --debug sync
crules --debug agent select
```

### Check Configuration

View current configuration settings:

```bash
cat ~/.config/crules/config.json
```

### Check Registry

View registered projects:

```bash
cat ~/.local/share/crules/registry.json
```

### Log Examination

Check log files for detailed error information:

```bash
# View most recent log file
ls -la ~/.local/share/crules/logs/ | sort | tail -n 1
cat ~/.local/share/crules/logs/latest.log
```

## Complete Reset

If you want to completely reset crules to its initial state:

```bash
# Remove configuration
rm -rf ~/.config/crules/

# Remove data files
rm -rf ~/.local/share/crules/

# Remove logs
rm -rf ~/.local/share/crules/logs/

# Run a simple command to reinitialize
crules --version
```

## Getting Help

If you're still experiencing issues:

1. Check for open issues on GitHub: [github.com/org/crules/issues](https://github.com/org/crules/issues)
2. File a new issue with:
   - Your operating system and version
   - Output of `crules --version`
   - Full error message and context
   - Steps to reproduce the issue

## Common Error Messages and Meanings

| Error Message | Likely Cause | Solution |
|---------------|--------------|----------|
| "Failed to create config directory" | Permission issues or disk space | Check directory permissions and available disk space |
| "Main rules location does not exist" | First time setup or incorrect path | Follow setup prompts or update config manually |
| "Failed to parse config file" | Corrupted configuration file | Reset configuration file |
| "Agent not found" | Missing agent files or typo in agent ID | Verify agent exists, check spelling, run `crules agent` to list available agents |
| "Failed to sync rules" | Permission issues or connectivity problems | Check file permissions, ensure main location is accessible |
| "Failed to register project" | Registry file issues | Check registry file permissions |
| "Failed to copy rules" | File system issues | Check directory permissions and disk space |

## See Also

- [Getting Started Guide](./getting-started.md): Quick introduction to crules
- [Configuration Guide](./configuration.md): How to configure crules
- [Command Reference](./commands.md): Detailed information about all commands

## Still Having Issues?

If you're still encountering problems after trying the solutions in this guide:

1. Check our [GitHub issues](https://github.com/nsnarender5511/crules/issues) to see if others have reported similar problems
2. Open a new issue with detailed information about your problem, including:
   - Your operating system and version
   - Your crules version (run `crules --version`)
   - Steps to reproduce the issue
   - Any error messages or logs
   - What you've already tried to resolve the issue

## Navigation

- Previous: [Agent System](./agents.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
