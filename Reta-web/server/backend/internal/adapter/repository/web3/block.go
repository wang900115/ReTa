package repositoryweb3

import (
	entitiesweb3 "backend/internal/domain/entities/web3"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type BlockRepository struct {
	DB *leveldb.DB
}

func NewBlockRepository(db *leveldb.DB) *BlockRepository {
	return &BlockRepository{DB: db}
}

func (repo *BlockRepository) GetLatestBlock() (*entitiesweb3.Block, error) {
	iter := repo.DB.NewIterator(util.BytesPrefix([]byte("block_")), nil)
	defer iter.Release()
	var latestBlock *entitiesweb3.Block
	for iter.Next() {
		var block entitiesweb3.Block
		err := json.Unmarshal(iter.Value(), &block)
		if err != nil {
			return nil, err
		}
		if latestBlock == nil || block.Index > latestBlock.Index {
			latestBlock = &block
		}
	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	return latestBlock, nil
}

func (repo *BlockRepository) GetBlockByHash(hash string) (*entitiesweb3.Block, error) {
	data, err := repo.DB.Get([]byte("block_"+hash), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	var block entitiesweb3.Block
	err = json.Unmarshal(data, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (repo *BlockRepository) GetBlockByIndex(index int) (*entitiesweb3.Block, error) {
	data, err := repo.DB.Get([]byte("block_"+strconv.Itoa(index)), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	var block entitiesweb3.Block
	err = json.Unmarshal(data, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (repo *BlockRepository) GetBlockByTimestamp(timestamp int64) (*entitiesweb3.Block, error) {
	data, err := repo.DB.Get([]byte("block_"+strconv.FormatInt(timestamp, 10)), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	var block entitiesweb3.Block
	err = json.Unmarshal(data, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (repo *BlockRepository) SaveBlock(block *entitiesweb3.Block) error {
	blockData, err := json.Marshal(block)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("block_"+block.Hash), blockData, nil)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("block_"+strconv.Itoa(block.Index)), blockData, nil)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("block_"+strconv.FormatInt(block.Timestamp, 10)), blockData, nil)
	if err != nil {
		return err
	}
	return nil
}
