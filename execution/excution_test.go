package execution

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/test"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
)

func TestExcution_ExcuteTx(t *testing.T) {
	testAcc := test.GetTestAccount()
	testDataList := test.GetTestTxsList(testAcc)
	exec := NewExcution()

	for i, testdata := range testDataList {
		memDb := database.NewMemDb(100000)
		BuildTestAccount(memDb)
		db := database.NewCacheDb(10000, memDb)

		receipt := exec.ExcuteTx(testdata.Txs, db)
		if receipt.String() != testdata.Receipt.String() {
			fmt.Println("test index: ", i)
			fmt.Println("result: ", receipt.String())
			fmt.Println("expect: ", testdata.Receipt.String())
			panic("unexpect testdata")
		}

		cacheDbBalance := db.Get(types.SystemAddress)
		if !cacheDbBalance.Amount.EqualTo(testdata.Receipt.TotalFee) {
			fmt.Println("test index: ", i)
			fmt.Println("balance result: ", cacheDbBalance.Amount)
			fmt.Println("balance expect: ", testdata.Receipt.TotalFee)
			panic("unexpect testdata")
		}
	}
}

func BuildTestAccount(db database.DB) {
	accList := test.GetTestAccount()
	for _, acc := range accList {
		db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}
