package types

type TxsData struct {
	TxSlice  Txs
	TotalFee Amount
}

func NewTxsData(txs Txs, totalFee Amount) TxsData {
	return TxsData{
		TxSlice:  txs,
		TotalFee: totalFee,
	}
}
