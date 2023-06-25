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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db postgresql.Database, group *gin.RouterGroup) {
	tr := infrastructure.NewTaskInfrastructure(db, domain.TaskTable)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
