package authorization

import (
	"go_iam/pkg/log"

	authzv1 "go_iam/internal/pkg/model/authz/v1"

	"github.com/ory/ladon"
)

// Authorizer implement the authorize interface that use local repository to
// authorize the subject access review.
type Authorizer struct {
	warden ladon.Warden
}

// NewAuthorizer creates a local repository authorizer and returns it.
func NewAuthorizer(authorizationClient AuthorizationInterface) *Authorizer {
	return &Authorizer{
		warden: &ladon.Ladon{
			Manager:     NewPolicyManager(authorizationClient),
			AuditLogger: NewAuditLogger(authorizationClient),
		},
	}
}

// Authorize to determine the subject access.
func (a *Authorizer) Authorize(request *ladon.Request) *authzv1.Response {
	log.Debug("authorize request", log.Any("request", request))

	if err := a.warden.IsAllowed(request); err != nil {
		return &authzv1.Response{
			Denied: true,
			Reason: err.Error(),
		}
	}

	return &authzv1.Response{
		Allowed: true,
	}
}
