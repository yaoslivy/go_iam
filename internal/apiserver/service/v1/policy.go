package v1

import (
	"context"
	"go_iam/internal/apiserver/store"
	"go_iam/internal/pkg/code"
	v1 "go_iam/internal/pkg/model/apiserver/v1"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

type PolicySrv interface {
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
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

func (s *policyService) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	if err := s.store.Policies().Create(ctx, policy, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

var _ PolicySrv = (*policyService)(nil)

func newPolicies(srv *service) *policyService {
	return &policyService{store: srv.store}
}
