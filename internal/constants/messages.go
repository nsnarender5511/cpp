package constants

// Error messages
const (
	// General errors
	ErrFailedOperation = "Operation failed: %s"
	ErrCommandFailed   = "Command failed: %w"

	// Path and security errors
	ErrMainLocationMissing    = "Main rules location does not exist: %s"
	ErrPathSecurityViolation  = "Security warning: path is outside agents directory: %s"
	ErrFailedGetAbsPath       = "Failed to get absolute path: %v"
	ErrFailedGetAgentsDirPath = "Failed to get absolute path for local agents directory: %v"
	ErrInvalidAgentDefPath    = "Invalid agent definition path"

	// File operation errors
	ErrFailedLoadAgentContent = "Failed to load agent content: %w"
	ErrFailedGetCurrentDir    = "Cannot get current directory: %v"

	// Agent related messages
	MsgNoAgentsFound      = "No agents found."
	MsgAgentListCompleted = "Agent list completed successfully | agent_count=%d"
	MsgInitCompleted      = "Init command completed successfully"

	// UI messages
	MsgNextStepsHeader = "Next Steps:"
	MsgCreateEmptyRepo = "Create empty rules directory"
	MsgCloneFromGit    = "Clone from Git repository"
	MsgSelectSource    = "Select source for rules:"
	MsgEnterRepoURL    = "Enter Git repository URL:"
)

// Success messages
const (
	MsgInitSuccess = "Agents initialized successfully"
)
