package constants

var (
	True = true
)

type ContextKey string

const (
	ContextKeyTraceId       ContextKey = "traceId"
	ContextKeyRepository    ContextKey = "repository"
	UserContextKey          ContextKey = "userId"
	ContextKeyGridId        ContextKey = "gridId"
	DemateAccountContextKey ContextKey = "demateAccountId"
)
