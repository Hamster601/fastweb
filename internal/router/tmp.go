package router

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/opentracing/opentracing-go"
//	"github.com/uber/jaeger-client-go/config"
//	"github.com/uber/jaeger-lib/metrics"
//)
//
//func main() {
//	// 初始化Jaeger配置
//	cfg, _ := config.FromEnv()
//
//	// 初始化Jaeger 实例和跟踪器
//	tracer, closer, _ := cfg.NewTracer(
//		config.Logger(jaeger.StdLogger),
//		config.Metrics(metrics.NullFactory),
//	)
//
//	// 延迟关闭 tracer
//	defer closer.Close()
//
//	// 注册 Jaeger 为全局跟踪器
//	opentracing.SetGlobalTracer(tracer)
//
//	// 初始化gin引擎
//	r := gin.Default()
//
//	// 定义一个处理器
//	r.GET("/hello", func(c *gin.Context) {
//		// 创建子跟踪, 并通过 gin的Context进行关联
//		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "hello")
//		defer span.Finish()
//
//		// 执行其他操作
//		c.JSON(200, gin.H{
//			"message": "hello world",
//		})
//
//	})
//
//	// 开始运行Gin引擎
//	r.Run(":8080")
//}

// elasticsearch
//import (
//	"go.uber.org/zap"
//	"go.elastic.co/apm/module/apmzap"
//	"github.com/kivutar/zap-elasticsearch-hook"
//)

//func main() {
//	// 创建 zap 日志对象
//	logger, err := zap.NewProduction()
//	logger = logger.WithOptions(zap.Hooks(apmzap.NewZapHook()))
//	if err != nil {
//		panic(err)
//	}
//
//	// 创建 elasticsearch 的钩子
//	elasticsearchHook, err := zapelasticsearch.NewElasticsearchHook(
//		"http://localhost:9200",
//		zapelasticsearch.Index("my-log-index"),
//		zapelasticsearch.Username("username"),
//		zapelasticsearch.Password("password"),
//	)
//	if err != nil {
//		logger.Error("Failed to create Elasticsearch hook", zap.Error(err))
//	}
//	defer elasticsearchHook.Close()
//
//	// 添加到 logger 钩子中
//	logger = logger.WithOptions(zap.Hooks(elasticsearchHook.Fire))
//
//	// 定义一些API路由
//	router := gin.Default()
//	router.GET("/api/users", func(c *gin.Context) {
//		logger.Info("Get all the users")
//		c.JSON(200, gin.H{
//			"status": "ok",
//		})
//	})
//	router.POST("/api/users", func(c *gin.Context) {
//		logger.Info("Post a new user")
//		c.JSON(201, gin.H{
//			"status": "created",
//		})
//	})
//
//	// 启动 gin 引擎
//	err = router.Run(":8080")
//	if err != nil {
//		panic(err)
//	}
//}
