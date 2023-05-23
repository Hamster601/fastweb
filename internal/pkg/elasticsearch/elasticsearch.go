package elasticsearch

// 参考文章
// https://www.jianshu.com/p/db4595fed6ea
// https://studygolang.com/articles/18494

//import (
//	"github.com/elastic/go-elasticsearch/v8"
//	"github.com/gin-gonic/gin"
//	zapelasticsearch "github.com/kivutar/zap-elasticsearch-hook"
//	"go.uber.org/zap"
//)
//
//func main() {
//	// 创建 zap 日志对象
//	logger, err := zap.NewProduction()
//	client, err := elasticsearch.NewDefaultClient()
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
