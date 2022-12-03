package policy

import (
	"go_iam/internal/pkg/middleware"
	"go_iam/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// Delete deletes the policy by the policy identifier.
func (p *PolicyController) Delete(c *gin.Context) {
	log.L(c).Info("delete policy function called.")

	username := c.GetString(middleware.UsernameKey)
	if err := p.srv.Policies().Delete(c, username, c.Param("name"), metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
