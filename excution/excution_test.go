package excution

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/mock"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
)

func TestExcution_ExcuteTx(t *testing.T) {
	testAcc := mock.GetTestAccount()
	exec := NewExcution()

	testDataList := []struct {
		txs     types.Txs
		receipt types.Receipt
	}{
		{
			txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(1)}},
			receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: mock.InitUserAmount.Sub(2), ToBalance: mock.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: mock.InitUserAmount, Fee: types.ToAmount(1)}},
			receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: false, FromBalance: mock.InitUserAmount, ToBalance: mock.InitUserAmount, SystemBalance: types.ZeroAmount}},
				TotalFee: types.ToAmount(0),
			},
		},
	}

	for i, testdata := range testDataList {
		ctx := mock.NewContext()
		receipt := exec.ExcuteTx(testdata.txs, ctx.Db)
		if receipt.String() != testdata.receipt.String() {
			fmt.Println("test index: ", i)
			fmt.Println("result: ", receipt.String())
			fmt.Println("expect: ", testdata.receipt.String())
			panic("unexpect testdata")
		}
	}
}
