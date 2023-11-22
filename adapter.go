package sqlcadapter

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	PolicyKey           = "casbin:policy"
	defaultTableName    = "casbin_rule"
	defaultDatabaseName = "casbin"
)

type CasbinRule struct {
	PType string `json:"ptype"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

type Adapter struct {
	cache    sqlc.CachedConn
	redisCli *redis.Client
	//然后再自己持有一个go-redis的客户端

}

func NewAdapter(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *Adapter {
	client := redis.NewClient(&redis.Options{
		Addr:     c[0].Host,
		Password: c[0].Pass,
		DB:       0,
	})
	newConn := sqlc.NewConn(conn, c)
	return &Adapter{
		newConn,
		client,
	}
}
