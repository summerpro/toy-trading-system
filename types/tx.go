package types

import "encoding/json"

type Txs []Tx

type Tx struct {
	From   Address
	To     Address
	Amount Amount
	Fee    Amount
}

func (tx *Txs) Serialize() []byte {
	res, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return res
}

func (tx *Txs) UnSerialize(jsonTx []byte) error {
	err := json.Unmarshal(jsonTx, tx)
	return err
}

func UnSerializeTxs(jsonTxs []byte) (txs Txs, err error) {
	err = json.Unmarshal(jsonTxs, &txs)
	return txs, err
}

func (tx *Tx) Validate() bool {
	if tx.Fee.Validate() == false {
		return false
	}
	if tx.Amount.Validate() == false {
		return false
	}
	if tx.From.Validate() == false {
		return false
	}
	if tx.To.Validate() == false {
		return false
	}
	return true
}
