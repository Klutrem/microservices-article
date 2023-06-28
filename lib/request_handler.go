package lib

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

type logger interface {
}

type GinLogger struct {
	*zap.SugaredLogger
}

func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

// NewRequestHandler creates a new request handler
func NewRequestHandler() RequestHandler {
	config := zap.NewDevelopmentConfig()
	level := zapcore.DebugLevel
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.SetLevel(level)
	zapLogger, _ := config.Build()
	log := zapLogger.Sugar().WithOptions(zap.WithCaller(false))
	logger := GinLogger{SugaredLogger: log}
	gin.DefaultWriter = logger
	engine := gin.New()
	return RequestHandler{Gin: engine}
}
