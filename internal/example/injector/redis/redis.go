package redis

import (
	"context"
	"fmt"

	"github.com/tangx/rum-gonic/rum"
)

/* 定义 */
type RedisAgent struct {
	Addr string
	Port int
}

func (r *RedisAgent) ServerAddr() string {
	return fmt.Sprintf("%s:%d", r.Addr, r.Port)
}

/* 初始化与注入 */
type RedisKey string

const (
	redisOnlineKey RedisKey = "online"
)

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
