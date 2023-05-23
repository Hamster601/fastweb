package router

import (
	"github.com/Hamster601/fastweb/internal/router/middlewares"
	"go.uber.org/zap"

	"github.com/Hamster601/fastweb/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	//Mux        core.Mux
	//Db    mysql.Repo
	//Cache redis.Repo
	//CronServer cron.Server
	GinEngine *gin.Engine
}

func NewHTTPServer(logger *zap.Logger, cronLogger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}
	defaultEngine := gin.Default()
	defaultEngine.Use(middlewares.LimitRate(), middlewares.Validator(), gin.Recovery())
	initAPIRouter(defaultEngine)
	s := new(Server)
	s.GinEngine = defaultEngine
	return s, nil
}
