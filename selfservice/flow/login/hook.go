package login

import (
	"github.com/ory/kratos/selfservice/flow"
)

type HookExecutorProvider interface {
	LoginHookExecutor() *flow.HookExecutor
}
