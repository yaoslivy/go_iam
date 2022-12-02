package secret

import (
	"go_iam/internal/pkg/middleware"
	"go_iam/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

func (s *SecretController) Delete(c *gin.Context) {
	log.L(c).Info("delete secret function called.")
	opts := metav1.DeleteOptions{Unscoped: true}
	username := c.GetString(middleware.UsernameKey)

	if err := s.srv.Secrets().Delete(c, username, c.Param("name"), opts); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
