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

func initAccount(accList []types.Account, db database.DB) {
	for _, acc := range accList {
		db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}
