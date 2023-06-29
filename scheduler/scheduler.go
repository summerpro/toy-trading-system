package scheduler

import (
	"github.com/summerpro/toy-trading-system/config"
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/execution"
	"github.com/summerpro/toy-trading-system/txpool"
	"github.com/summerpro/toy-trading-system/types"
	"log"
	"time"
)

type Scheduler struct {
	txPool   *txpool.HeapTxPool
	db       database.DB
	exec     *execution.Execution
	txsChan  chan types.Txs
	stopChan chan struct{}
	context  *config.Context
}

func NewSchdeduler(context *config.Context) *Scheduler {
	return &Scheduler{
		txPool:   txpool.NewHeapTxPool(context.InitTxsPoolSize, context.MaxTxsPoolSize),
		db:       database.NewMemDb(context.InitMemDbSize),
		exec:     execution.NewExcution(),
		txsChan:  make(chan types.Txs, context.TxsChannelSize),
		stopChan: make(chan struct{}),
		context:  context,
	}
}

func (scheduler *Scheduler) Schedule() {
	go scheduler.SolveTxs()

	tick := time.Tick(time.Millisecond * time.Duration(scheduler.context.ExecTxSleepTime))
	for {
		select {
		case <-scheduler.stopChan:
			return
		case <-tick:
			if scheduler.txPool.Size() > 0 {
				txsData := scheduler.txPool.GetTopTx(scheduler.context.BlockSize)
				cacheDb := database.NewCacheDb(scheduler.context.InitCacheDbSize, scheduler.db)
				receiptList := make([]types.Receipt, 0, len(txsData))
				for _, txs := range txsData {
					receipt := scheduler.exec.ExcuteTx(txs.TxSlice, cacheDb)
					receiptList = append(receiptList, receipt)
				}
				cacheDb.Commit()
				cacheDb.RangeCache()
			}

		}
	}

}

func (scheduler *Scheduler) ReceiveTxs(txs types.Txs) {
	scheduler.txsChan <- txs
}

func (scheduler *Scheduler) SolveTxs() {
	for txs := range scheduler.txsChan {
		for scheduler.txPool.Size() == scheduler.context.MaxTxsPoolSize {
			time.Sleep(time.Millisecond * time.Duration(scheduler.context.ExecTxSleepTime))
		}
		cacheDb := database.NewCacheDb(scheduler.context.InitCacheDbSize, scheduler.db)
		receipt := scheduler.exec.ExcuteTx(txs, cacheDb)

		if receipt.TotalFee.LargerThan(types.ZeroAmount) {
			txsData := types.TxsData{
				TxSlice:  txs,
				TotalFee: receipt.TotalFee,
			}
			scheduler.txPool.InsertTx(txsData)
			log.Println("valid txs: ", string(txs.Serialize()))
		} else {
			log.Println("invalid txs: ", string(txs.Serialize()))
		}

	}
}

func (scheduler *Scheduler) StopScheduler() {
	scheduler.stopChan <- struct{}{}
}
