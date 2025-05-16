# Vibe CLI

Vibe is a command-line tool designed to manage and interact with AI agents, potentially for an IDE or development environment. It allows users to initialize agent configurations, list available agents, select agents for use, and view detailed information about them.

## Installation / Running

Currently, Vibe can be run directly from its source code using Go:

```bash
# Make sure you are in the root directory of the vibe project
go run ./cmd <command> [options]
```

For example:
```bash
go run ./cmd init
go run ./cmd agent list
```

*(Build instructions for creating a standalone binary can be added here if available in the future).*

## Usage

The general command-line syntax for Vibe is:

```bash
vibe [GLOBAL OPTIONS] <command> [COMMAND OPTIONS] [ARGUMENTS...]
```
Or when running with `go run`:
```bash
go run ./cmd [GLOBAL OPTIONS] <command> [COMMAND OPTIONS] [ARGUMENTS...]
```

## Global Options

These options can be used with any command:

*   `--debug`: Show detailed debug, informational, and error messages. This is primarily for development or troubleshooting and provides maximum output detail.
*   `--multi-agent`: Enable multi-agent mode for the current session. This updates the configuration to reflect that multiple agents might be active or coordinated.
*   `-v`: Display the current version of Vibe and exit.

## Commands

Vibe offers the following main commands:

### `init`

Initializes the current directory to work with Vibe agents.

**Usage:**
```bash
vibe init
```
Or:
```bash
go run ./cmd init
```

**What it does:**
1.  **Agent System Initialization**: Sets up the necessary files and directories for Vibe agents (typically within a `.cursor/` subdirectory). This may include deploying default agent rule files (`.mdc` files).
2.  **Ignore File Updates**:
    *   Adds `.cursor/` to the project's `.gitignore` file.
    *   Adds `.cursorignore` to the project's `.gitignore` file.
    *   Adds `.cursor/` to a `.cursorignore` file.
    This ensures that Vibe's operational files are not accidentally committed to version control or processed by Vibe itself in unwanted ways.
3.  **Post-Initialization Information**:
    *   Lists a few of the agents that are now available in the initialized directory.
    *   Provides "Next Steps" to guide the user on how to further interact with agents (e.g., `vibe agent list`).

**Common Use Pattern:**
Run `vibe init` once when you want to start using Vibe agents in a new project or directory.

```bash
# Navigate to your project directory
cd /path/to/your/project

# Initialize Vibe
go run ../path/to/vibe/cmd init
```

### `agent`

Manages and interacts with available agents. This command has several subcommands.

**Usage:**
```bash
vibe agent <subcommand> [ARGUMENTS...]
```
Or:
```bash
go run ./cmd agent <subcommand> [ARGUMENTS...]
```

#### `agent list`

Lists all agents available to Vibe in the current context (checking the local project and then system-wide agent definitions).

**Usage:**
```bash
vibe agent list
```
**Output:**
Displays a formatted list of agents, potentially grouped by category. The agent last selected (if any) might be highlighted. The display adapts to terminal width.

**Example:**
```bash
go run ./cmd agent list
```

#### `agent select`

Allows you to interactively select an agent from the list of available agents. Once selected, the agent is loaded and its ID is saved as the "last selected agent" for future reference (e.g., highlighting in `agent list`).

**Usage:**
```bash
vibe agent select
```
**Interaction:**
Presents a navigable list of agents. After selection, it confirms the loaded agent and asks if you want to view its details.

**Common Use Pattern:**
Use this when you want to choose an agent to work with and see its details or make it the default for subsequent (conceptual) operations.

```bash
go run ./cmd agent select
```

#### `agent info <agent_name_or_id_or_index>`

Displays detailed information about a specific agent. The agent can be identified by its unique ID, its name, or its 1-based index from the `agent list` output.

**Usage:**
```bash
vibe agent info <identifier>
```
*   `<identifier>`: The name, ID, or 1-based index of the agent.

**Output:**
Shows detailed information about the agent, including its ID, Name, Version (if any), Type (if any), Last Updated timestamp, Description, and Templates. If the global `--verbose` (or `--debug`) flag is used, it will also show the agent's specific configuration parameters.

**Examples:**
```bash
go run ./cmd agent info my-custom-agent
go run ./cmd agent info architect-planner
go run ./cmd agent info 3
go run ./cmd --verbose agent info wizard
```

#### `agent run <agent_name_or_id_or_index>`

(Currently Conceptual) Runs a specific agent. The agent can be identified by its unique ID, name, or 1-based index.

**Usage:**
```bash
vibe agent run <identifier>
```
*   `<identifier>`: The name, ID, or 1-based index of the agent.

**Current Behavior (as of refactoring):**
This command will find the specified agent and print a message indicating it's a "Conceptual run." The full execution logic for agents is not yet implemented in this command path. It does update the "last selected agent" configuration.

**Future Behavior (Intended):**
This command would load and then execute the specified agent's primary task or function.

**Examples:**
```bash
go run ./cmd agent run my-custom-agent
go run ./cmd agent run 1
```

#### `agent help` (or `--help`, `-h`)

Displays the help message specifically for the `agent` command, showing its subcommands and examples.

**Usage:**
```bash
vibe agent help
```

## How Options Influence Behavior

*   **`--debug`**: This flag controls the log level via `internal/utils/logger.go` (setting `debugConsole`, `verbose`, and `verboseErrors` states). `utils.Debug()`, `utils.Info()` messages will be printed. `cmd/cli/utils.go`'s `HandleCommandError` function will print detailed errors. The `agent info` command also uses this state (passed as `verbose` to its handler) to show detailed agent configuration.
*   **`--multi-agent`**: When this flag is used, `cmd/main.go` loads the configuration via `internal/utils/config_manager.go`, sets the `MultiAgentEnabled` field to `true`, and saves the configuration. This setting might be checked by agent execution logic in the future.
*   **`-v`**: `cmd/main.go` calls `internal/version/version.go`'s `GetVersion()` function and prints the result.

## Common Usage Patterns

1.  **Setting up a new project:**
    ```bash
    cd my-new-project
    go run /path/to/vibe/cmd init
    ```

2.  **Listing available agents:**
    ```bash
    go run /path/to/vibe/cmd agent list
    ```

3.  **Getting help for a command:**
    ```bash
    go run /path/to/vibe/cmd --help         # General help
    go run /path/to/vibe/cmd agent --help  # Help for the 'agent' command
    ```

4.  **Viewing details of a specific agent (e.g., "wizard"):**
    ```bash
    go run /path/to/vibe/cmd agent info wizard
    ```
    For even more details, including its configuration (now enabled by `--debug`):
    ```bash
    go run /path/to/vibe/cmd --debug agent info wizard
    ```

5.  **Interactively selecting an agent:**
    ```bash
    go run /path/to/vibe/cmd agent select
    ```

## Project Structure Overview

*   **`cmd/main.go`**: Main application entry point. Handles global flag setup and command dispatching.
*   **`cmd/cli/`**: Contains all CLI-specific logic:
    *   `flags.go`: Defines global command-line flags.
    *   `ui.go`: Functions to print usage and help messages.
    *   `commands.go`: Dispatches commands to their respective handlers.
    *   `init_cmd.go`: Handler for the `init` command.
    *   `agent_cmd.go`: Handler for the `agent` command and its subcommands.
    *   `utils.go`: CLI-specific utility functions (e.g., `HandleCommandError`).
*   **`internal/`**: Houses the core libraries and business logic of Vibe:
    *   `agent/`: Agent definition, loading, and registry management.
    *   `core/`: Core system components like `AgentInitializer`.
    *   `ui/`: Reusable UI components for console output (banners, lists, formatted text).
    *   `utils/`: General utility functions (config management, logging, file operations).
    *   `version/`: Application version management.
    *   `constants/`: (Assumed) For global constants.
    *   `git/`: (Assumed) For Git-related functionalities.

---

This README provides a starting point. It can be expanded further with more details on agent creation, configuration file structure, or advanced usage scenarios as the Vibe project evolves. 