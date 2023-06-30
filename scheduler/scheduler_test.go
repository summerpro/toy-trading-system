package scheduler

import (
	"github.com/summerpro/toy-trading-system/config"
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/testdata"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
	"time"
)

func TestScheduler_Schedule(t *testing.T) {
	sch := NewSchdeduler(config.DefaultContext())
	testAcc := testdata.GetTestAccount()
	testDataList := testdata.GetTestTxsList(testAcc)
	initAccount(testAcc, sch.db)

	sch.StartScheduler()

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
	testAcc := testdata.GetTestAccount()
	testDataList := testdata.GetTestTxsList(testAcc)
	initAccount(testAcc, sch.db)

	sch.StartScheduler()

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
	testAcc := testdata.GetTestAccount()
	testDataList := []testdata.TestTxsData{
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(3)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(4)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(5)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
	}
	initAccount(testAcc, sch.db)

	sch.StartScheduler()

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
