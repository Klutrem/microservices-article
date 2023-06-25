package route

import (
	"time"

	"main/api/controller"
	"main/bootstrap"
	"main/domain"
	"main/infrastructure"
	"main/mongo"
	"main/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := infrastructure.NewUserInfrastructure(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
