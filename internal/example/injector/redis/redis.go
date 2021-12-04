package redis

import (
	"context"

	"github.com/tangx/rum-gonic/rum"
)

type RedisKey string

const (
	redisOnlineKey RedisKey = "online"
)

type RedisAgent struct {
	Addr string
	Port int
}

var RedisOnlineInjector = rum.NewContextInjector(
	redisOnlineKey,
	&RedisAgent{
		Addr: "Online-agent",
		Port: 6379,
	},
)

func FromRedisAgentOnline(ctx context.Context) *RedisAgent {
	key := rum.WithContextKey(redisOnlineKey)
	val := ctx.Value(key)

	return val.(*RedisAgent)
}
