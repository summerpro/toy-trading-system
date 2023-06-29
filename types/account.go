package types

import "encoding/json"

type Account struct {
	Addr   Address
	Amount Amount
}

func (account *Account) Serialize() []byte {
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	return res
}

func (account *Account) UnSerialize(jsonAccount []byte) error {
	err := json.Unmarshal(jsonAccount, account)
	return err
}

func UnSerializeAccount(jsonAccount []byte) (acc Account) {
	err := json.Unmarshal(jsonAccount, &acc)
	if err != nil {
		panic(err)
	}
	return acc
}
