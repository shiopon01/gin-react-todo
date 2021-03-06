package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// init ...
// Connect to DB using `engine`
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "gin:password@/gin_sample")
	if err != nil {
		panic(err)
	}
}
