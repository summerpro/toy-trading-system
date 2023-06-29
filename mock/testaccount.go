package mock

import (
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/types"
)

const (
	InitUserAmount types.Amount = 1000
)

func BuildTestAccount(db database.DB) {
	accList := GetTestAccount()
	for _, acc := range accList {
		db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}

func GetTestAccount() []types.Account {
	return []types.Account{
		{
			Addr:   "alice",
			Amount: InitUserAmount,
		},
		{
			Addr:   "bob",
			Amount: InitUserAmount,
		},
		{
			Addr:   "tom",
			Amount: InitUserAmount,
		},
		{
			Addr:   "ray",
			Amount: InitUserAmount,
		},
	}
}

func GetEmptyAccountList() []types.Account {
	return []types.Account{
		{
			Addr:   "emptyUser1",
			Amount: types.ZeroAmount,
		},
		{
			Addr:   "emptyUser2",
			Amount: types.ZeroAmount,
		},
		{
			Addr:   "emptyUser3",
			Amount: types.ZeroAmount,
		},
		{
			Addr:   "emptyUser4",
			Amount: types.ZeroAmount,
		},
	}
}
