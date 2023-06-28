package txpool

import (
	"github.com/summerpro/toy-trading-system/types"
	"sync"
)

type TxPool struct {
	size     int
	hp       interface{}
	poolLock sync.Mutex
}

func (txPool *TxPool) InsertTx(tx *types.Tx) {

}

func (txPool *TxPool) GetTopTx(txNum int) []types.Tx {
	return nil
}
