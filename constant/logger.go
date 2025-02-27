package constant

const (
	Level   = "level"
	Time    = "time"
	Message = "message"
	Value   = "value"

	TimeFormat = "2006/01/02 - 15:04:05"

	RenderResponseMessage  = "middlewareHandler.RenderResponse executed and response rendered"
	RequestReceivedMessage = "route.helper.BuildAndSetRequestMetaInContext executed and requestlogged"

	LogLevelPanic = "panic"
	LogLevelFatal = "fatal"
	LogLevelError = "error"
	LogLevelWarn  = "warning"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
	LogLevelTrace = "trace"
)
