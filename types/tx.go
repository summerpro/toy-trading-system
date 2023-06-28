package types

import "encoding/json"

type Tx []struct {
	From   string
	To     string
	Amount int64
	Fee    int64
}

func (tx *Tx) Serialize() []byte {
	res, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return res
}

func (tx *Tx) UnSerialize(jsonTx []byte) error {
	err := json.Unmarshal(jsonTx, tx)
	return err
}
