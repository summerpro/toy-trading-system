package excution

import (
	"github.com/summerpro/toy-trading-system/database"
	"github.com/summerpro/toy-trading-system/types"
)

type Excution struct {
}

func (e *Excution) ExcuteTx(tx types.Tx, db database.DB) {

}

func (e *Excution) ExcuteTxWithoutCommit(tx types.Tx, db database.DB) {

}
