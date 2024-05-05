package application

import (
	"main/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

// TestController data type
type Controller struct {
	logger pkg.Logger
}

// NewTestController creates new Test controller
func NewController(logger pkg.Logger) Controller {
	return Controller{
		logger: logger,
	}
}

func (u Controller) Test1(c *gin.Context) {
	time.Sleep(time.Second)

	u.logger.Debug("getting request on test1")

	c.JSON(200, gin.H{
		"message": "ok",
		"status":  200,
	})

}

func (u Controller) Test2(c *gin.Context) {
	time.Sleep(time.Second * 2)

	u.logger.Debug("getting request on test2")

	c.JSON(200, gin.H{
		"message": "ok",
		"status":  200,
	})

}

func (u Controller) Test3(c *gin.Context) {
	time.Sleep(time.Second * 3)

	u.logger.Debug("getting request on test3")

	c.JSON(200, gin.H{
		"message": "ok",
		"status":  200,
	})

}
