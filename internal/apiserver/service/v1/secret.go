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

var _ SecretSrv = (*secretService)(nil)

func newSecrets(srv *service) *secretService {
	return &secretService{store: srv.store}
}
