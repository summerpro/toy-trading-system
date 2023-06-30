package testdata

import (
	"github.com/summerpro/toy-trading-system/types"
)

const (
	InitUserAmount types.Amount = 100000000
)

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

type TestTxsData struct {
	Txs     types.Txs
	Receipt types.Receipt
}

func GetTestTxsList(testAcc []types.Account) []TestTxsData {
	testDataList := []TestTxsData{
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(1)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: InitUserAmount.Sub(2), ToBalance: InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: InitUserAmount, Fee: types.ToAmount(1)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: false, FromBalance: InitUserAmount, ToBalance: InitUserAmount, SystemBalance: types.ZeroAmount}},
				TotalFee: types.ToAmount(0),
			},
		},
	}
	return testDataList
}
