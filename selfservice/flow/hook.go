package flow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/session"
)

type (
	PreHookExecutor interface {
		ExecutePreFlowHook(w http.ResponseWriter, r *http.Request, a Flow) error
	}
	PreHookExecutorFunc func(w http.ResponseWriter, r *http.Request, a Flow) error

	PostHookPostPersistExecutor interface {
		ExecutePostFlowPostPersistHook(w http.ResponseWriter, r *http.Request, a Flow, s *session.Session) error
	}
	PostHookPostPersistExecutorFunc func(w http.ResponseWriter, r *http.Request, a Flow, s *session.Session) error

	PostHookPrePersistExecutor interface {
		ExecutePostFlowPrePersistHook(w http.ResponseWriter, r *http.Request, a Flow, i *identity.Identity) error
	}
	PostHookPrePersistExecutorFunc func(w http.ResponseWriter, r *http.Request, a Flow, i *identity.Identity) error

	HooksProvider interface {
		PreFlowHooks(ctx context.Context) []PreHookExecutor
		PostFlowPrePersistHooks(ctx context.Context, credentialsType identity.CredentialsType) []PostHookPrePersistExecutor
		PostFlowPostPersistHooks(ctx context.Context, credentialsType identity.CredentialsType) []PostHookPostPersistExecutor
	}
)

func PostHookPostPersistExecutorNames(e []PostHookPostPersistExecutor) []string {
	names := make([]string, len(e))
	for k, ee := range e {
		names[k] = fmt.Sprintf("%T", ee)
	}
	return names
}

func (f PreHookExecutorFunc) ExecutePreFlowHook(w http.ResponseWriter, r *http.Request, a Flow) error {
	return f(w, r, a)
}
func (f PostHookPostPersistExecutorFunc) ExecutePostFlowPostPersistHook(w http.ResponseWriter, r *http.Request, a Flow, s *session.Session) error {
	return f(w, r, a, s)
}
func (f PostHookPrePersistExecutorFunc) ExecutePostFlowPrePersistHook(w http.ResponseWriter, r *http.Request, a Flow, i *identity.Identity) error {
	return f(w, r, a, i)
}

type (
	HookExecutor struct {
		d HooksProvider
	}
)

func NewHookExecutor(d HooksProvider) *HookExecutor {
	return &HookExecutor{d: d}
}

func (e *HookExecutor) PostFlowPrePersistHook(w http.ResponseWriter, r *http.Request, ct identity.CredentialsType, a Flow, i *identity.Identity) error {
	for _, e := range e.d.PostFlowPrePersistHooks(r.Context(), ct) {
		if err := e.ExecutePostFlowPrePersistHook(w, r, a, i); err != nil {
			return err
		}
	}

	return nil
}

func (e *HookExecutor) PostFlowPostPersistHook(w http.ResponseWriter, r *http.Request, ct identity.CredentialsType, a Flow, s *session.Session) error {
	for _, e := range e.d.PostFlowPostPersistHooks(r.Context(), ct) {
		if err := e.ExecutePostFlowPostPersistHook(w, r, a, s); err != nil {
			return err
		}
	}

	return nil
}

func (e *HookExecutor) PreFlowHook(w http.ResponseWriter, r *http.Request, a Flow) error {
	for _, executor := range e.d.PreFlowHooks(r.Context()) {
		if err := executor.ExecutePreFlowHook(w, r, a); err != nil {
			return err
		}
	}

	return nil
}
