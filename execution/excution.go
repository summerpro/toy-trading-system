package execution

import (
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/types"
)

type Execution struct {
}

func NewExcution() *Execution {
	return &Execution{}
}

func (e *Execution) ExcuteTx(txs types.Txs, db *database.CacheDb) types.Receipt {
	receipt := types.NewReceipt(len(txs))
	for _, tx := range txs {
		var fromAccount, toAccount types.Account
		txInvalidFlag := false

		systemAccount := getAccount(types.SystemAddress, db)
		if tx.Validate() {
			fromAccount = getAccount(tx.From, db)
			toAccount = getAccount(tx.To, db)
		} else {
			txInvalidFlag = true
		}

		var receiptItem types.ReceiptItem
		if validateTx(tx, fromAccount, toAccount) == false || txInvalidFlag {
			receiptItem.Access = false
		} else {
			receiptItem.Access = true
			fromAccount.Amount = fromAccount.Amount.Sub(tx.Amount.Add(tx.Fee))
			toAccount.Amount = toAccount.Amount.Add(tx.Amount)
			systemAccount.Amount = systemAccount.Amount.Add(tx.Fee)
			receipt.TotalFee = receipt.TotalFee.Add(tx.Fee)
			db.Set(tx.From, fromAccount)
			db.Set(tx.To, toAccount)
		}
		receiptItem.FromBalance = fromAccount.Amount
		receiptItem.ToBalance = toAccount.Amount
		receiptItem.SystemBalance = systemAccount.Amount

		db.Set(types.SystemAddress, systemAccount)
		receipt.Item = append(receipt.Item, receiptItem)
	}
	return receipt
}

func getAccount(addr types.Address, db *database.CacheDb) types.Account {
	account := db.Get(addr)
	return account
}

func validateTx(tx types.Tx, fromAccount types.Account, toAccount types.Account) bool {
	if fromAccount.Amount.LessThan(tx.Amount.Add(tx.Fee)) {
		return false
	}
	if tx.From == tx.To {
		return false
	}
	return true
}
