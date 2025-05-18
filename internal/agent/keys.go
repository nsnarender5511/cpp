package agent


type ContextKey string


const (
	KeyAgentID        ContextKey = "agent_id"
	KeyAgentType      ContextKey = "agent_type"
	KeyAgentVersion   ContextKey = "agent_version"
	KeyLastExecution  ContextKey = "last_execution"
	KeyExecutionCount ContextKey = "execution_count"
	KeyErrorCount     ContextKey = "error_count"
	KeyMetadata       ContextKey = "metadata"
	KeyCustomData     ContextKey = "custom_data"
)
