package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/miladhzzzz/milx-cloud-init/api-gateway/controllers"
)

type GithubRouteController struct {
	authController   controllers.AuthController
	githubController controllers.GithubController
}

func NewGithubRouteController(githubController controllers.GithubController) GithubRouteController {
	return GithubRouteController{githubController: githubController}
}

func (rc *GithubRouteController) GithubRoute(rg *gin.RouterGroup) {
	router := rg.Group("")

	router.POST("/webhook", rc.githubController.WebhookHandler())

	private := router.Group("github")

	private.Use(rc.authController.Auth())

	private.GET("/set/webhook", rc.githubController.SetWebhook())
	private.GET("/set/accessToken", rc.githubController.SetAccessToken())

}
