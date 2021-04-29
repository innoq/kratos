package hook

import (
	"net/http"

	"github.com/ory/kratos/selfservice/flow"

	"github.com/ory/kratos/session"
)

type (
	sessionDestroyerDependencies interface {
		session.ManagementProvider
		session.PersistenceProvider
	}
	SessionDestroyer struct {
		r sessionDestroyerDependencies
	}
)

func NewSessionDestroyer(r sessionDestroyerDependencies) *SessionDestroyer {
	return &SessionDestroyer{r: r}
}

func (e *SessionDestroyer) ExecutePostFlowPostPersistHook(_ http.ResponseWriter, r *http.Request, _ flow.Flow, s *session.Session) error {
	if err := e.r.SessionPersister().DeleteSessionsByIdentity(r.Context(), s.Identity.ID); err != nil {
		return err
	}
	return nil
}
