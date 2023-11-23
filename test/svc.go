package test

import (
	"fmt"
	sqlcadapter "github.com/peifengll/sqlc-adapter"
	"github.com/peifengll/sqlc-adapter/config"
	"github.com/peifengll/sqlc-adapter/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ada *sqlcadapter.Adapter
var casrule model.CasbinRuleModel

func Init() {
	c := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DbName,
	)
	mysqlConn := sqlx.NewMysql(dsn)
	ada = sqlcadapter.NewAdapter(mysqlConn, c.Cache)
	casrule = model.NewCasbinRuleModel(mysqlConn, c.Cache)
}
