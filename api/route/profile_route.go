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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := infrastructure.NewUserInfrastructure(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
