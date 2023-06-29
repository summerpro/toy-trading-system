package txpool

import (
	"container/heap"
	"github.com/summerpro/toy-trading-system/types"
	"sync"
)

type Heap []types.TxsData

func (heap Heap) Len() int {
	return len(heap)
}

func (heap Heap) Less(i, j int) bool {
	return heap[i].TotalFee.LargerThan(heap[j].TotalFee)
}

func (heap Heap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *Heap) Push(data interface{}) {
	*heap = append(*heap, data.(types.TxsData))
}

func (heap *Heap) Pop() interface{} {
	old := *heap
	length := len(old)
	x := old[length-1]
	*heap = old[0 : length-1]
	return x
}

type TxsHeap struct {
	txsHeap *Heap
	lock    sync.Mutex
}

func NewTxsHeap(size int) *TxsHeap {
	txsHeap := make(Heap, 0, size)
	heap.Init(&txsHeap)
	return &TxsHeap{txsHeap: &txsHeap}
}

func (txsHeap *TxsHeap) Push(txs types.TxsData) {
	txsHeap.Lock()
	defer txsHeap.UnLock()

	heap.Push(txsHeap.txsHeap, txs)
}

func (txsHeap *TxsHeap) Pop() types.TxsData {
	txsHeap.Lock()
	defer txsHeap.UnLock()

	return heap.Pop(txsHeap.txsHeap).(types.TxsData)
}

func (txsHeap *TxsHeap) Size() int {
	txsHeap.Lock()
	defer txsHeap.UnLock()

	return len(*txsHeap.txsHeap)
}

func (txsHeap *TxsHeap) Lock() {
	txsHeap.lock.Lock()
}

func (txsHeap *TxsHeap) UnLock() {
	txsHeap.lock.Unlock()
}