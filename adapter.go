package sqlcadapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
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
	Id    int64  `db:"id"`
	Ptype string `db:"ptype"`
	V0    string `db:"v0"`
	V1    string `db:"v1"`
	V2    string `db:"v2"`
	V3    string `db:"v3"`
	V4    string `db:"v4"`
	V5    string `db:"v5"`
}
type Adapter struct {
	cache    sqlc.CachedConn
	redisCli *redis.Client
	table    string
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
		defaultTableName,
	}
}

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(model model.Model) (err error) {
	ctx := context.Background()
	// Using the LoadPolicyLine handler from the Casbin repo for building rules
	return a.loadPolicy(ctx, model, persist.LoadPolicyArray)
}

func (a *Adapter) loadPolicy(ctx context.Context, model model.Model, handler func([]string, model.Model) error) (err error) {
	// 0, -1 fetches all entries from the list
	//rules, err := a.redisCli.LRange(ctx, PolicyKey, 0, -1).Result()
	//if err != nil {
	//	return err
	//}
	//
	//// Parse the rules from Redis
	//for _, rule := range rules {
	//	handler(strings.Split(rule, ", "), model)
	//}
	return
}

// Adapter is the interface for Casbin adapters.
type Adapte interface {
	// LoadPolicy loads all policy rules from the storage.
	LoadPolicy(model model.Model) error
	// SavePolicy saves all policy rules to the storage.
	SavePolicy(model model.Model) error

	// AddPolicy adds a policy rule to the storage.
	// This is part of the Auto-Save feature.
	AddPolicy(sec string, ptype string, rule []string) error
	// RemovePolicy removes a policy rule from the storage.
	// This is part of the Auto-Save feature.
	RemovePolicy(sec string, ptype string, rule []string) error
	// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
	// This is part of the Auto-Save feature.
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
}

func (a *Adapter) SelectRows() ([]*CasbinRule, error) {
	policyKey := PolicyKey
	resp := make([]*CasbinRule, 100)
	ctx := context.Background()
	//a.cache.QueryRowNoCache()
	//a.cache.QueryRowsNoCache()
	err := a.cache.QueryRowCtx(ctx, resp, policyKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select * from %s", a.table)
		return conn.QueryRowCtx(ctx, v, query)
	})
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}

func (a *Adapter) SelectRow() (*CasbinRule, error) {
	policyKey := PolicyKey
	var resp CasbinRule
	ctx := context.Background()
	err := a.cache.QueryRowCtx(ctx, &resp, policyKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select * from %s limit 1", a.table)
		return conn.QueryRowCtx(ctx, v, query)
	})

	//var resp CasbinRule
	//err := a.cache.QueryRowCtx(ctx, &resp, policyKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
	//	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", "*", a.table)
	//	return conn.QueryRowCtx(ctx, v, query, 1)
	//})

	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}
}
