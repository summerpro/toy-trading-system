package txpool

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
)

func TestHeapTxPool(t *testing.T) {
	txSlice := []types.TxsData{
		{TotalFee: types.ToAmount(1)},
		{TotalFee: types.ToAmount(3)},
		{TotalFee: types.ToAmount(5)},
		{TotalFee: types.ToAmount(20)},
		{TotalFee: types.ToAmount(34)},
		{TotalFee: types.ToAmount(2)},
		{TotalFee: types.ToAmount(23)},
		{TotalFee: types.ToAmount(6)},
		{TotalFee: types.ToAmount(3)},
	}
	newHeapPool := NewHeapTxPool(100000, 100000)
	for _, txs := range txSlice {
		newHeapPool.InsertTx(txs)
	}

	maxTxs := newHeapPool.GetTopTx(5)
	fmt.Println(maxTxs)

	if newHeapPool.txsHeap.Size() != 4 {
		panic("unexpected HeapPool size")
	}
}
