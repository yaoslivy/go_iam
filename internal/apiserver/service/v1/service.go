package v1

import "go_iam/internal/apiserver/store"

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
}

type service struct {
	store store.Factory
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}

func (s *service) Policies() PolicySrv {
	return newPolicies(s)
}

// NewService returns Service interface
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
