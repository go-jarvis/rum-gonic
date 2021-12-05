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

var RedisOnlineAgent = &RedisAgent{
	Addr: "Online-agent",
	Port: 6379,
}

func FromRedisAgentOnline(ctx context.Context) *RedisAgent {
	val := rum.FromContextValue(ctx, redisOnlineKey)

	return val.(*RedisAgent)
}

func WithRedisInject() rum.ContextInjectorFunc {
	return rum.WithContextValue(redisOnlineKey, RedisOnlineAgent)
}
