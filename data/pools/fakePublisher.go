package pools

import (
	"fmt"

	algodtxqueue "github.com/Diatomix/algod-tx-queue"
	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/data/transactions"
)

type FakePublisher struct {
	cfg config.Local
}

func (f *FakePublisher) StreamTxGroup(txGroup []transactions.SignedTxn) {
	fmt.Printf("I will broadcast a transaction, here is my redis env: %s\n", f.cfg.RedisDSN)
	algodtxqueue.MustStreamTxGroup(txGroup)
}
