// Command goAi is the malatd service project.
// The framework reference: https://github.com/swxctx/malatd
package __TPL__

// __API_TPL__ register PULL router
type __API_TPL__ interface {
	V1_Test
}

type V1_Test interface {
	Ping(*PingArgsV1) *PingResultV1
}

type (
	PingArgsV1   = struct{}
	PingResultV1 = struct{}
)
