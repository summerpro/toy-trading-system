package types

import "encoding/json"

type Receipt struct {
	Item     []ReceiptItem
	TotalFee Amount
}

type ReceiptItem struct {
	Access        bool
	FromBalance   Amount
	ToBalance     Amount
	SystemBalance Amount
}

func NewReceipt(num int) Receipt {
	item := make([]ReceiptItem, 0, num)
	return Receipt{
		Item:     item,
		TotalFee: ZeroAmount,
	}
}

func (receipt *Receipt) String() string {
	res, err := json.Marshal(receipt)
	if err != nil {
		panic(err)
	}
	return string(res)
}
