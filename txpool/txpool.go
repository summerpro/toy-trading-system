package txpool

import (
	"github.com/summerpro/toy-trading-system/types"
	"sync"
)

type HeapTxPool struct {
	maxSize  int
	txsHeap  *TxsHeap
	poolLock sync.Mutex
}

func NewHeapTxPool(maxSize int, initPoolSize int) *HeapTxPool {
	return &HeapTxPool{
		maxSize:  maxSize,
		txsHeap:  NewTxsHeap(initPoolSize),
		poolLock: sync.Mutex{},
	}
}

func (txPool *HeapTxPool) InsertTx(txs types.TxsData) {
	txPool.Lock()
	defer txPool.UnLock()

	if txPool.txsHeap.Size() < txPool.maxSize {
		txPool.txsHeap.Push(txs)
	}
}

func (txPool *HeapTxPool) GetTopTx(txNum int) []types.TxsData {
	txPool.Lock()
	defer txPool.UnLock()

	txsList := make([]types.TxsData, 0, txNum)

	for i := 0; i < txNum; i++ {
		if txPool.txsHeap.Size() == 0 {
			break
		}
		txs := txPool.txsHeap.Pop()
		txsList = append(txsList, txs)
	}
	return txsList
}

func (txPool *HeapTxPool) Lock() {
	txPool.poolLock.Lock()
}

func (txPool *HeapTxPool) UnLock() {
	txPool.poolLock.Unlock()
}

func (txPool *HeapTxPool) Size() int {
	txPool.Lock()
	defer txPool.UnLock()

	return txPool.txsHeap.Size()
}
