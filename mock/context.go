package mock

import "github.com/summerpro/toy-trading-system/database"

type Context struct {
	Db database.DB
}

func NewContext() *Context {
	db := database.NewMemDb()
	BuildTestAccount(db)

	return &Context{
		Db: db,
	}
}
