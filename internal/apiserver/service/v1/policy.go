package v1

import (
	"context"
	"go_iam/internal/apiserver/store"
	v1 "go_iam/internal/pkg/model/apiserver/v1"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type PolicySrv interface {
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type policyService struct {
	store store.Factory
}

func (s *policyService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	policies, err := s.store.Policies().List(ctx, username, opts)
	if err != nil {
		return nil, err
	}
	return policies, nil
}

var _ PolicySrv = (*policyService)(nil)

func newPolicies(srv *service) *policyService {
	return &policyService{store: srv.store}
}
