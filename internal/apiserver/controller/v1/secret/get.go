package secret

import (
	"go_iam/internal/pkg/middleware"
	"go_iam/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// Get get an policy by the secret identifier.
func (s *SecretController) Get(c *gin.Context) {
	log.L(c).Info("get secret function called.")

	username := c.GetString(middleware.UsernameKey)
	secret, err := s.srv.Secrets().Get(c, username, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, secret)
}
