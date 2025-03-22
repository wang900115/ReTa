package repositoryweb3

import (
	entitiesweb3 "backend/internal/domain/entities/web3"

	"github.com/goccy/go-json"
	"github.com/syndtr/goleveldb/leveldb"
)

type BlockchainRepository struct {
	DB *leveldb.DB
}

func NewBlockchainRepository(db *leveldb.DB) *BlockchainRepository {
	return &BlockchainRepository{DB: db}
}

func (repo *BlockchainRepository) GetBlockchainHeight() (int, error) {
	data, err := repo.DB.Get([]byte("blockchain_height"), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	var height int
	err = json.Unmarshal(data, &height)
	if err != nil {
		return 0, err
	}
	return height, nil
}

func (repo *BlockchainRepository) SaveBlockchainState(bc *entitiesweb3.Blockchain) error {
	heightData, err := json.Marshal(bc.Height)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("blockchain_height"), heightData, nil)
	if err != nil {
		return err
	}

	bcData, err := json.Marshal(bc)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("blockchain_state"), bcData, nil)
	if err != nil {
		return err
	}

	return nil
}
