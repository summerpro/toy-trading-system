package execution

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/testdata"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
)

func TestExcution_ExcuteTx(t *testing.T) {
	testAcc := testdata.GetTestAccount()
	testDataList := testdata.GetTestTxsList(testAcc)
	exec := NewExcution(10)

	for i, testdata := range testDataList {
		memDb := database.NewMemDb(100000)
		BuildTestAccount(memDb)
		db := database.NewCacheDb(10000, memDb)

		receipt := exec.ExcuteTx(testdata.Txs, db)
		if receipt.String() != testdata.Receipt.String() {
			fmt.Println("testdata index: ", i)
			fmt.Println("result: ", receipt.String())
			fmt.Println("expect: ", testdata.Receipt.String())
			panic("unexpect testdata")
		}

		cacheDbBalance := db.Get(types.SystemAddress)
		if !cacheDbBalance.Amount.EqualTo(testdata.Receipt.TotalFee) {
			fmt.Println("testdata index: ", i)
			fmt.Println("balance result: ", cacheDbBalance.Amount)
			fmt.Println("balance expect: ", testdata.Receipt.TotalFee)
			panic("unexpect testdata")
		}
	}
}

func BuildTestAccount(db database.DB) {
	accList := testdata.GetTestAccount()
	for _, acc := range accList {
		db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}
