package mock

import "github.com/summerpro/toy-trading-system/database"

const (
	initMemDbSize   = 100000
	initCacheDbSize = 10000
)

type Context struct {
	Db database.DB
}

func NewContext() *Context {
	db := database.NewMemDb(initMemDbSize)
	BuildTestAccount(db)

	return &Context{
		Db: db,
	}
}
