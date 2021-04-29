package driver

import (
	"context"

	"github.com/ory/kratos/selfservice/flow"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/login"
)

func (m *RegistryDefault) LoginHookExecutor() *flow.HookExecutor {
	if m.selfserviceLoginExecutor == nil {
		m.selfserviceLoginExecutor = flow.NewHookExecutor(&loginHooksProvider{m: m})
	}
	return m.selfserviceLoginExecutor
}

type loginHooksProvider struct {
	m *RegistryDefault
}

func (lhp *loginHooksProvider) PreFlowHooks(ctx context.Context) (b []flow.PreHookExecutor) {
	for _, v := range lhp.m.getHooks("", lhp.m.Config(ctx).SelfServiceFlowLoginBeforeHooks()) {
		if hook, ok := v.(flow.PreHookExecutor); ok {
			b = append(b, hook)
		}
	}
	return
}

func (lhp *loginHooksProvider) PostFlowPrePersistHooks(_ context.Context, _ identity.CredentialsType) (b []flow.PostHookPrePersistExecutor) {
	return
}

func (lhp *loginHooksProvider) PostFlowPostPersistHooks(ctx context.Context, credentialsType identity.CredentialsType) (b []flow.PostHookPostPersistExecutor) {
	for _, v := range lhp.m.getHooks(string(credentialsType), lhp.m.Config(ctx).SelfServiceFlowLoginAfterHooks(string(credentialsType))) {
		if hook, ok := v.(flow.PostHookPostPersistExecutor); ok {
			b = append(b, hook)
		}
	}

	if len(b) == 0 {
		// since we don't want merging hooks defined in a specific strategy and global hooks
		// global hooks are added only if no strategy specific hooks are defined
		for _, v := range lhp.m.getHooks("global", lhp.m.Config(ctx).SelfServiceFlowLoginAfterHooks("global")) {
			if hook, ok := v.(flow.PostHookPostPersistExecutor); ok {
				b = append(b, hook)
			}
		}
	}
	return
}

func (m *RegistryDefault) LoginHandler() *login.Handler {
	if m.selfserviceLoginHandler == nil {
		m.selfserviceLoginHandler = login.NewHandler(m)
	}

	return m.selfserviceLoginHandler
}

func (m *RegistryDefault) LoginFlowErrorHandler() *login.ErrorHandler {
	if m.selfserviceLoginRequestErrorHandler == nil {
		m.selfserviceLoginRequestErrorHandler = login.NewFlowErrorHandler(m)
	}

	return m.selfserviceLoginRequestErrorHandler
}
