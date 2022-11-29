package v1

import (
	"context"
	"go_iam/internal/apiserver/store"
	"go_iam/internal/pkg/code"
	v1 "go_iam/internal/pkg/model/apiserver/v1"
	"regexp"

	"github.com/marmotedu/errors"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

//UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
}

type userService struct {
	store store.Factory
}

func (u *userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	if err := u.store.Users().Create(ctx, user, opts); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error()); match {
			return errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}
