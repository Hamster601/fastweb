package resolvers

import (
	"context"
	"github.com/Hamster601/fastweb/internal/graph/generated"
	"github.com/Hamster601/fastweb/internal/pkg/core"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/mysql"
	"github.com/Hamster601/fastweb/internal/pkg/infraDB/redis"
	"go.uber.org/zap"
)

type coreCtxKeyType struct{ name string }

var CoreContextKey = coreCtxKeyType{"_core_context"}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type Resolver struct {
	logger *zap.Logger
	cache  redis.Repo
	//userService user_service.UserService
}

func NewRootResolvers(logger *zap.Logger, db mysql.Repo, cache redis.Repo) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			logger: logger,
			cache:  cache,
			//userService: user_service.NewUserService(db, cache),
		},
	}
	return c
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// getCoreContextByCtx 获取 core context
func (r *Resolver) getCoreContextByCtx(ctx context.Context) core.Context {
	return ctx.Value(CoreContextKey).(core.Context)
}
