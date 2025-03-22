package entitiesweb3

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Blockchain struct {
	Height      int
	Difficulty  int
	LatestBlock *Block
	DB          *leveldb.DB
}
