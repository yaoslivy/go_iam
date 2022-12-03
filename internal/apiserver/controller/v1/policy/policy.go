package policy

import (
	srvv1 "go_iam/internal/apiserver/service/v1"
	"go_iam/internal/apiserver/store"
)

// PolicyController create a policy handler used to handle request for policy resource.
type PolicyController struct {
	srv srvv1.Service
}

//
func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{
		srv: srvv1.NewService(store),
	}
}
