package irepositoryweb3

import entitiesweb3 "backend/internal/domain/entities/web3"

type IBlockchainRepository interface {
	GetBlockchainHeight() (*entitiesweb3.Blockchain, error) // Get blockchain height
	SaveBlockchainState(bc *entitiesweb3.Blockchain) error  // Get blockchain status
}
