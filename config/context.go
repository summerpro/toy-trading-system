package config

type Context struct {
	InitCacheDbSize int
	InitMemDbSize   int
	InitTxsPoolSize int
	ExecTxSleepTime int // ms
	BlockSize       int
	MaxTxsPoolSize  int
	TxsChannelSize  int
}

func NewContext(initCacheDbSize int,
	initMemDbSize int,
	execTxSleepTime int,
	initTxsPoolSize int,
	blockSize int,
	maxTxsPoolSize int,
	txsChannelSize int) *Context {
	return &Context{
		InitCacheDbSize: initCacheDbSize,
		InitMemDbSize:   initMemDbSize,
		InitTxsPoolSize: initTxsPoolSize,
		ExecTxSleepTime: execTxSleepTime,
		BlockSize:       blockSize,
		MaxTxsPoolSize:  maxTxsPoolSize,
		TxsChannelSize:  txsChannelSize,
	}
}

func DefaultContext() *Context {
	return &Context{
		InitCacheDbSize: DefaultInitCacheDbSize,
		InitMemDbSize:   DefaultInitMemDbSize,
		InitTxsPoolSize: DefaultInitTxsPoolSize,
		ExecTxSleepTime: DefaultExecTxSleepTime,
		BlockSize:       DefaultBlockSize,
		MaxTxsPoolSize:  DefaultMaxTxsPoolSize,
		TxsChannelSize:  DefaultTxsChannelSize,
	}
}
