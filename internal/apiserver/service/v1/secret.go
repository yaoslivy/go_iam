package v1

import (
	"context"
	"go_iam/internal/apiserver/store"
	"go_iam/internal/pkg/code"
	v1 "go_iam/internal/pkg/model/apiserver/v1"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

type SecretSrv interface {
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error)
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error
	Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error)
}

type secretService struct {
	store store.Factory
}

func (s *secretService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.SecretList, error) {
	secrets, err := s.store.Secrets().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return secrets, nil
}

func (s *secretService) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) error {
	if err := s.store.Secrets().Create(ctx, secret, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (s *secretService) Get(ctx context.Context, username, secretID string, opts metav1.GetOptions) (*v1.Secret, error) {
	secret, err := s.store.Secrets().Get(ctx, username, secretID, opts)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

var _ SecretSrv = (*secretService)(nil)

func newSecrets(srv *service) *secretService {
	return &secretService{store: srv.store}
}
