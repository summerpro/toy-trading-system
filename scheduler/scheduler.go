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
		exec:     execution.NewExcution(context.MaxTxsSize),
		txsChan:  make(chan types.Txs, context.TxsChannelSize),
		stopChan: make(chan struct{}),
		context:  context,
	}
}

func (scheduler *Scheduler) schedule() {
	go scheduler.solveTxs()

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
				log.Println("time to solve txs,", len(txsData), "txs have been solved")
				cacheDb.Commit()
				cacheDb.RangeCache()
			}

		}
	}

}

func (scheduler *Scheduler) ReceiveTxs(txs types.Txs) {
	scheduler.txsChan <- txs
}

func (scheduler *Scheduler) solveTxs() {
	for txs := range scheduler.txsChan {
		poolSize := scheduler.txPool.Size()
		for poolSize == scheduler.context.MaxTxsPoolSize {
			log.Println("txs pool is full, size: ", poolSize)
			time.Sleep(time.Millisecond * time.Duration(scheduler.context.ExecTxSleepTime))
		}
		log.Println("txs pool size: ", poolSize)
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
	close(scheduler.txsChan)
	close(scheduler.stopChan)
}

func (scheduler *Scheduler) StartScheduler() {
	go scheduler.schedule()
}

func (scheduler *Scheduler) InitAccount(initAccount []types.Account) {
	for _, acc := range initAccount {
		scheduler.db.Set(acc.Addr.Bytes(), acc.Serialize())
	}
}
