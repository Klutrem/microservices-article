package route

import (
	"time"

	"main/api/controller"
	"main/bootstrap"
	"main/domain"
	"main/infrastructure"
	"main/pkg/postgresql"
	"main/usecase"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db postgresql.Database, group *gin.RouterGroup) {
	ur := infrastructure.NewUserInfrastructure(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}