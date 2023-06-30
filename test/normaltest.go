package main

import (
	"github.com/summerpro/toy-trading-system/config"
	"github.com/summerpro/toy-trading-system/scheduler"
	"github.com/summerpro/toy-trading-system/testdata"
	"github.com/summerpro/toy-trading-system/types"
	"time"
)

func main() {
	cfg := config.DefaultContext()
	cfg.BlockSize = 5
	cfg.ExecTxSleepTime = 1000
	sch := scheduler.NewSchdeduler(cfg)
	testAcc := testdata.GetTestAccount()
	testDataList := buildTestDataList(testAcc)
	sch.InitAccount(testAcc)

	sch.StartScheduler()

	for {
		for _, testData := range testDataList {
			go sch.ReceiveTxs(testData.Txs)
		}

		time.Sleep(time.Millisecond * 3000)
	}
}

func buildTestDataList(testAcc []types.Account) []testdata.TestTxsData {
	testDataList := []testdata.TestTxsData{
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(-1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(3)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(0)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(4)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(0)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(5)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(0)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(2)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(-1), Fee: types.ToAmount(0)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(1), Fee: types.ToAmount(10)}}},
		{Txs: []types.Tx{{From: testAcc[0].Addr, To: testAcc[1].Addr, Amount: types.ToAmount(-1), Fee: types.ToAmount(10)}, {From: testAcc[2].Addr, To: testAcc[3].Addr, Amount: types.ToAmount(-1), Fee: types.ToAmount(10)}}},
	}
	return testDataList
}
