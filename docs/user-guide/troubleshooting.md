# Troubleshooting Guide

> 🔧 This guide helps you diagnose and resolve common issues that may occur when using cursor++.

## Quick Reference

### Common Issues by Category

#### Installation Issues
- Command `cursor++` not found error
- "Command not found" when trying to run cursor++
- Permission denied errors

#### Installation Troubleshooting

**Symptoms:**
- Command not found error
- Unable to execute cursor++

**Solutions:**
1. Check if cursor++ is installed and in your PATH:
   ```bash
   which cursor++
   ```

2. Check your PATH environment variable:
   ```bash
   echo $PATH
   ```

3. Reinstall cursor++:
   ```bash
   curl -sL https://github.com/org/cursor++/releases/latest/download/install.sh | bash
   ```

4. Add cursor++ to your PATH:
   ```bash
   # Add to ~/.bashrc or ~/.zshrc
   export PATH=$PATH:/path/to/cursor++/bin
   # Then reload
   source ~/.bashrc # or source ~/.zshrc
   ```

#### Permission Issues

**Symptoms:**
- "Permission denied" when running cursor++

**Solutions:**
1. Make the binary executable:
   ```bash
   chmod +x /path/to/cursor++
   ```

2. Check file permissions:
   ```bash
   ls -la /path/to/cursor++
   ```

3. Try running with sudo:
   ```bash
   sudo cursor++ --version
   ```

### Configuration Issues

#### Config File Not Found

**Symptoms:**
- "Config file not found" error
- "Failed to load configuration"

**Solutions:**
1. Check if config directory exists:
   ```bash
   ls -la ~/.config/cursor++/
   ```

2. Create a default config file:
   ```json
   // ~/.config/cursor++/config.json
   {
     "mainLocation": "~/.cursor/rules"
   }
   ```

3. Try initializing again:
   ```bash
   cursor++ init
   ```

#### Permission Errors with Config

**Symptoms:**
- "Permission denied" when accessing config file
- Unable to write to config directory

**Solutions:**
1. Check directory permissions:
   ```bash
   ls -la ~/.config/
   ```

2. Create the directory with proper permissions:
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
   sudo cursor++ init
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
   cursor++ --debug sync
   ```

4. Try resetting by reinitializing:
   ```bash
   cursor++ init
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
   cursor++ list
   ```

3. Check for file conflicts:
   ```bash
   cursor++ --debug merge
   ```

4. Try with `--verbose` flag:
   ```bash
   cursor++ --verbose merge
   ```

### Registry Issues

#### Projects Missing from Registry

**Symptoms:**
- Projects not showing in `cursor++ list`
- "Project not registered" errors

**Solutions:**
1. Re-register the project:
   ```bash
   cursor++ init
   ```

2. Check registry file for corruption:
   ```bash
   cat ~/.local/share/cursor++/registry.json
   ```

3. Reset registry file:
   ```bash
   rm ~/.local/share/cursor++/registry.json
   ```

4. Run registry repair:
   ```bash
   cursor++ clean
   ```

#### Cannot Clean Non-existent Projects

**Symptoms:**
- Error when running `cursor++ clean`
- Non-existent projects still showing in list

**Solutions:**
1. Check registry file permissions:
   ```bash
   ls -la ~/.local/share/cursor++/registry.json
   ```

2. Manually edit registry file:
   ```bash
   # Edit registry file directly
   nano ~/.local/share/cursor++/registry.json
   ```

3. Reset registry completely:
   ```bash
   echo '{"projects":[]}' > ~/.local/share/cursor++/registry.json
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
   cursor++ sync
   ```

4. Import agents from URL:
   ```bash
   cursor++ import https://cursor.directory/example-agent
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
   cursor++ --debug init
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
   cursor++ --verbose init
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
- Error when running `cursor++ agent select`
- Selection UI crashes or exits unexpectedly

**Solutions:**
1. Try selecting by ID directly:
   ```bash
   cursor++ agent info wizard
   ```

2. Check terminal compatibility:
   ```bash
   echo $TERM
   ```

3. Update terminal settings:
   ```bash
   export TERM=xterm-256color
   ```

4. Try running with debug flags:
   ```bash
   cursor++ --debug agent select
   ```

### Configuration File Issues

#### Corrupt Configuration

**Symptoms:**
- "Failed to parse config" error
- "Invalid configuration format"

**Solutions:**
1. Check the current configuration:
   ```bash
   cat ~/.config/cursor++/config.json
   ```

2. Reset the configuration file:
   ```bash
   rm ~/.config/cursor++/config.json
   cursor++ --version  # Creates new default config
   ```

3. Create a minimal configuration file:
   ```bash
   echo '{"mainLocation":"~/.cursor/rules"}' > ~/.config/cursor++/config.json
   ```

### Permission and Ownership Issues

#### Permission Denied Errors

**Symptoms:**
- "Permission denied" when accessing configuration
- Unable to create or modify rules files

**Solutions:**
1. Check file permissions:
   ```bash
   ls -la ~/.config/cursor++/
   ls -la ~/.local/share/cursor++/
   ```

2. Fix permissions:
   ```bash
   chmod 755 ~/.config/cursor++/
   chmod 644 ~/.config/cursor++/config.json
   chmod 755 ~/.local/share/cursor++/
   chmod 644 ~/.local/share/cursor++/registry.json
   ```

3. Fix ownership:
   ```bash
   chown -R $(whoami) ~/.config/cursor++/
   chown -R $(whoami) ~/.local/share/cursor++/
   ```

## Common Commands for Diagnosis

### Check Installation Status

Check that cursor++ is properly installed and get version information:

```bash
cursor++ --version
```

### Debug Mode for Operations

```bash
cursor++ --debug init
cursor++ --debug sync
cursor++ --debug agent select
```

### Check Configuration Files

```bash
# View configuration file
cat ~/.config/cursor++/config.json

# View registry file
cat ~/.local/share/cursor++/registry.json
```

### View Logs

```bash
# List log files
ls -la ~/.local/share/cursor++/logs/ | sort | tail -n 1
cat ~/.local/share/cursor++/logs/latest.log
```

### Complete Reset

If you want to completely reset cursor++ to its initial state:

```bash
# Remove configuration
rm -rf ~/.config/cursor++/

# Remove registry
rm -rf ~/.local/share/cursor++/

# Remove logs
rm -rf ~/.local/share/cursor++/logs/

# Create fresh installation
cursor++ --version
```

## Getting Help

### Community Resources

1. Check for open issues on GitHub: [github.com/org/cursor++/issues](https://github.com/org/cursor++/issues)
2. When reporting an issue, include:
   - Output of `cursor++ --version`
   - Error messages you're seeing
   - Steps to reproduce the issue
   - Operating system and version

### Common Error Messages

| Error Message | Likely Cause | Solution |
|---------------|--------------|----------|
| "Command not found" | Installation problem | Add cursor++ to PATH or reinstall |
| "Failed to initialize" | Permission or configuration issue | Check permissions, recreate config |
| "Rules directory not found" | Directory structure issue | Create required directories |
| "Agent not found" | Missing agent files or typo in agent ID | Verify agent exists, check spelling, run `cursor++ agent` to list available agents |
| "Cannot sync rules" | Permission or path issue | Check directory permissions and paths |

## See Also

- [Getting Started Guide](./getting-started.md): Quick introduction to cursor++
- [Configuration Guide](./configuration.md): How to configure cursor++
- [Command Reference](./commands.md): Detailed information about all commands

## Still Having Issues?

If you're still encountering problems after trying the solutions in this guide:

1. Check our [GitHub issues](https://github.com/nsnarender5511/vibe/issues) to see if others have reported similar problems
2. Open a new issue with detailed information about your problem, including:
   - Your operating system and version
   - Your cursor++ version (run `cursor++ --version`)
   - Steps to reproduce the issue
   - Any error messages or logs
   - What you've already tried to resolve the issue

## Navigation

- Previous: [Agent System](./agents.md)
- Up: [User Guide](../README.md#user-guide)
- Home: [Documentation Home](../README.md)
