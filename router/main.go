package router

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func SetRouter(router *gin.Engine, buildFS embed.FS) {
	context := router.Group("/oapi")
	SetApiRouter(context)
	SetDashboardRouter(context)
	SetRelayRouter(context)
	frontendBaseUrl := os.Getenv("FRONTEND_BASE_URL")
	//if config.IsMasterNode && frontendBaseUrl != "" {
	//	frontendBaseUrl = ""
	//	logger.SysLog("FRONTEND_BASE_URL is ignored on master node")
	//}
	if frontendBaseUrl == "" {
		SetWebRouter(router, buildFS)
	} else {
		frontendBaseUrl = strings.TrimSuffix(frontendBaseUrl, "/")
		router.NoRoute(func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s%s", frontendBaseUrl, c.Request.RequestURI))
		})
	}
}
