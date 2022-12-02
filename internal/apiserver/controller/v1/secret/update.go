package secret

import (
	"go_iam/internal/pkg/code"
	"go_iam/internal/pkg/middleware"
	v1 "go_iam/internal/pkg/model/apiserver/v1"
	"go_iam/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
)

// Update update a key by the secret key identifier.
func (s *SecretController) Update(c *gin.Context) {
	log.L(c).Info("update secret function called.")

	var r v1.Secret

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	username := c.GetString(middleware.UsernameKey)
	name := c.Param("name")

	// only update expires and description
	secret, err := s.srv.Secrets().Get(c, username, name, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	// only update expires and description
	secret.Expires = r.Expires
	secret.Description = r.Description
	secret.Extend = r.Extend

	if errs := secret.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)
		return
	}

	if err := s.srv.Secrets().Update(c, secret, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, secret)
}
