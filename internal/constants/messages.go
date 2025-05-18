package constants


const (
	
	ErrFailedOperation = "Operation failed: %s"
	ErrCommandFailed   = "Command failed: %v"

	
	ErrMainLocationMissing    = "Main rules location does not exist: %s"
	ErrPathSecurityViolation  = "Security warning: path is outside agents directory: %s"
	ErrFailedGetAbsPath       = "Failed to get absolute path: %v"
	ErrFailedGetAgentsDirPath = "Failed to get absolute path for local agents directory: %v"
	ErrInvalidAgentDefPath    = "Invalid agent definition path"

	
	ErrFailedLoadAgentContent = "Failed to load agent content: %v"
	ErrFailedGetCurrentDir    = "Cannot get current directory: %v"

	
	MsgNoAgentsFound      = "No agents found."
	MsgAgentListCompleted = "Agent list completed successfully | agent_count=%d"
	MsgInitCompleted      = "Init command completed successfully"

	
	MsgNextStepsHeader = "Next Steps:"
	MsgCreateEmptyRepo = "Create empty rules directory"
	MsgCloneFromGit    = "Clone from Git repository"
	MsgSelectSource    = "Select source for rules:"
	MsgEnterRepoURL    = "Enter Git repository URL:"
)


const (
	MsgInitSuccess = "Agents initialized successfully"
)
