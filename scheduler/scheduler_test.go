package scheduler

import (
	"github.com/summerpro/toy-trading-system/config"
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/test"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
	"time"
)

func TestScheduler_Schedule(t *testing.T) {
	sch := NewSchdeduler(config.DefaultContext())
	testAcc := test.GetTestAccount()
	testDataList := test.GetTestTxsList(testAcc)
	initAccount(testAcc, sch.db)

	go sch.Schedule()

	for _, testData := range testDataList {
		sch.ReceiveTxs(testData.Txs)
	}
	time.Sleep(time.Millisecond * 100)

	for _, testData := range testDataList {
		sch.ReceiveTxs(testData.Txs)
	}
	time.Sleep(time.Millisecond * 100)

	sch.StopScheduler()

}

func TestScheduler_Schedule_MultiRequest(t *testing.T) {
	sch := NewSchdeduler(config.DefaultContext())
	testAcc := test.GetTestAccount()
	testDataList := test.GetTestTxsList(testAcc)
	initAccount(testAcc, sch.db)

	go sch.Schedule()

	for i := 0; i < 500; i++ {
		go func() {
			for _, testData := range testDataList {
				sch.ReceiveTxs(testData.Txs)
			}
		}()

	}

	time.Sleep(time.Millisecond * 1000)

	sch.StopScheduler()

}

func TestScheduler_Schedule_TopN(t *testing.T) {
	cfg := config.DefaultContext()
	cfg.BlockSize = 5
	cfg.ExecTxSleepTime = 1000
	sch := NewSchdeduler(cfg)
	testAcc := test.GetTestAccount()
	testDataList := []test.TestTxsData{
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(3)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(4)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(5)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
		{
			Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}},
			Receipt: types.Receipt{
				Item:     []types.ReceiptItem{{Access: true, FromBalance: test.InitUserAmount.Sub(2), ToBalance: test.InitUserAmount.Add(1), SystemBalance: types.ToAmount(1)}},
				TotalFee: types.ToAmount(1),
			},
		},
	}
	initAccount(testAcc, sch.db)

	go sch.Schedule()

	for _, testData := range testDataList {
		sch.ReceiveTxs(testData.Txs)
	}

	time.Sleep(time.Millisecond * 3000)

	sch.StopScheduler()

}

func initAccount(accList []types.Account, db database.DB) {
	for _, acc := range accList {
		db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}
