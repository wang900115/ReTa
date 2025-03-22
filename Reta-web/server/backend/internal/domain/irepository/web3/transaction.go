package irepositoryweb3

import entitiesweb3 "backend/internal/domain/entities/web3"

type ITranSactionRepository interface {
	GetTransactionByHash(hash string) (*entitiesweb3.Transaction, error)              // Get transaction by hash
	GetTransactionsByBlockHash(blockhash string) ([]*entitiesweb3.Transaction, error) // Get transaction by block's hash
	SaveTransaction(transaction *entitiesweb3.Transaction) error                      // Save transaction
}
