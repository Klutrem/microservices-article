package route

import (
	"time"

	"main/api/controller"
	"main/bootstrap"
	"main/domain"
	"main/mongo"
	"main/repository"
	"main/usecase"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
