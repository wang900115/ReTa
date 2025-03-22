package irepositoryweb3

import entitiesweb3 "backend/internal/domain/entities/web3"

type IBlockRepository interface {
	GetLatestBlock() (*entitiesweb3.Block, error)                     // Get latest block
	GetBlockByHash(hash string) (*entitiesweb3.Block, error)          // Get block by hash
	GetBlockByIndex(index int) (*entitiesweb3.Block, error)           // Get block by index
	GetBlockByTimestamp(timestamp int64) (*entitiesweb3.Block, error) // Get block by timestamp
	SaveBlock(block *entitiesweb3.Block) error                        // Save block
}
