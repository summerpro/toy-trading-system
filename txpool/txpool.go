package txpool

import (
	"github.com/summerpro/toy-trading-system/types"
	"sync"
)

type HeapTxPool struct {
	size     int
	hp       interface{}
	poolLock sync.Mutex
}

func (txPool *HeapTxPool) InsertTx(tx *types.Tx) {

}

func (txPool *HeapTxPool) GetTopTx(txNum int) []types.Tx {
	return nil
}
