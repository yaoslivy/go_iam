package v1

import (
	"context"
	"go_iam/internal/apiserver/store"
	"go_iam/internal/pkg/code"
	v1 "go_iam/internal/pkg/model/apiserver/v1"
	"go_iam/pkg/log"
	"regexp"
	"sync"

	"github.com/marmotedu/errors"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

//UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	ChangePassword(ctx context.Context, user *v1.User) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
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

func (u *userService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user, err := u.store.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	if err := u.store.Users().Update(ctx, user, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) ChangePassword(ctx context.Context, user *v1.User) error {
	// Save changed fields.
	if err := u.store.Users().Update(ctx, user, metav1.UpdateOptions{}); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *userService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if err := u.store.Users().Delete(ctx, username, opts); err != nil {
		return err
	}
	return nil
}

// List returns user list in the storage. This function has a good performance.
func (u *userService) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	users, err := u.store.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan error, 1)

	var m sync.Map

	// Improve query efficiency in parallel
	for _, user := range users.Items {
		wg.Add(1)

		go func(user *v1.User) {
			defer wg.Done()

			//some cost time process
			policies, err := u.store.Policies().List(ctx, user.Name, metav1.ListOptions{})
			if err != nil {
				errChan <- errors.WithCode(code.ErrDatabase, err.Error())
				return
			}

			m.Store(user.ID, &v1.User{
				ObjectMeta: metav1.ObjectMeta{
					ID:         user.ID,
					InstanceID: user.InstanceID,
					Name:       user.Name,
					Extend:     user.Extend,
					CreatedAt:  user.CreatedAt,
					UpdatedAt:  user.UpdatedAt,
				},
				Nickname:    user.Nickname,
				Email:       user.Email,
				Phone:       user.Phone,
				TotalPolicy: policies.TotalCount,
				LoginedAt:   user.LoginedAt,
			})
		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*v1.User, 0, len(users.Items))
	for _, user := range users.Items {
		info, _ := m.Load(user.ID)
		infos = append(infos, info.(*v1.User))
	}

	log.L(ctx).Debugf("get %d users from backend storage.", len(infos))

	return &v1.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}

func (u *userService) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {
	if err := u.store.Users().DeleteCollection(ctx, usernames, opts); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}
