package policy

import (
	"go_iam/internal/pkg/code"
	"go_iam/internal/pkg/middleware"
	"go_iam/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

// List returns all policies.
func (p *PolicyController) List(c *gin.Context) {
	log.L(c).Info("list policy function called.")

	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	username := c.GetString(middleware.UsernameKey)
	policies, err := p.srv.Policies().List(c, username, r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, policies)
}
