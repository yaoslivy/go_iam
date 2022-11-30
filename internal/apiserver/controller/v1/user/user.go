package user

import (
	srvv1 "go_iam/internal/apiserver/service/v1"
	"go_iam/internal/apiserver/store"
)

// UserController create a user handler used to handler request for user resource.
type UserController struct {
	srv srvv1.Service
}

// NewUserController creates a user handler.
func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
