package txpool

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/types"
	"testing"
)

func TestTxsHeap(t *testing.T) {
	txSlice := []types.TxsData{
		{TotalFee: types.ToAmount(1)},
		{TotalFee: types.ToAmount(3)},
		{TotalFee: types.ToAmount(5)},
		{TotalFee: types.ToAmount(7)},
		{TotalFee: types.ToAmount(34)},
		{TotalFee: types.ToAmount(2)},
		{TotalFee: types.ToAmount(5)},
		{TotalFee: types.ToAmount(6)},
		{TotalFee: types.ToAmount(3)},
	}
	newHeap := NewTxsHeap(100000)
	for _, txs := range txSlice {
		newHeap.Push(txs)
	}

	maxTxs := newHeap.Pop()
	if !maxTxs.TotalFee.EqualTo(types.ToAmount(34)) {
		fmt.Println("result: ", maxTxs.TotalFee)
		fmt.Println("expect: ", 34)
		panic("unexpect testdata")
	}
}
