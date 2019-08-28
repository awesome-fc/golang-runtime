package golangruntime

const (
	fcRequestID             = "x-fc-request-id"
	fcAccessKeyID           = "x-fc-access-key-id"
	fcAccessKeySecret       = "x-fc-access-key-secret"
	fcSecurityToken         = "x-fc-security-token"
	fcFunctionName          = "x-fc-function-name"
	fcFunctionHandler       = "x-fc-function-handler"
	fcFunctionMemory        = "x-fc-function-memory"
	fcFunctionTimeout       = "x-fc-function-timeout"
	fcFunctionInitializer   = "x-fc-function-initializer"
	fcInitializationTimeout = "x-fc-initialization-timeout"

	fcServiceName       = "x-fc-service-name"
	fcServiceLogProject = "x-fc-service-logproject"
	fcServiceLogstore   = "x-fc-service-logstore"

	fcRegion    = "x-fc-region"
	fcAccountID = "x-fc-account-id"
	fcQualifier = "x-fc-qualifier"
	fcVersionID = "x-fc-version-id"

	fcStatus      = "x-fc-status"
	fcControlPath = "x-fc-control-path"

	fcLogTailStartPrefix           = "FC Invoke Start RequestId: %s"     // Start of log tail mark
	fcLogTailEndPrefix             = "FC Invoke End RequestId: %s"       // End of log tail mark
	fcInitializeLogTailStartPrefix = "FC Initialize Start RequestId: %s" // Start of initialize log tail mark
	fcLogInitializeTailEndPrefix   = "FC Initialize End RequestId: %s"   // End of initialize log tail mark
)
