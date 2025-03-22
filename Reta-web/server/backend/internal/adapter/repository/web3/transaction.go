package repositoryweb3

import (
	entitiesweb3 "backend/internal/domain/entities/web3"
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type TransactionRepository struct {
	DB *leveldb.DB
}

func NewTransactionRepository(db *leveldb.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (repo *TransactionRepository) GetTransactionByHash(hash string) (*entitiesweb3.Transaction, error) {
	data, err := repo.DB.Get([]byte("transaction_"+hash), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	var transaction entitiesweb3.Transaction
	err = json.Unmarshal(data, &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (repo *TransactionRepository) GetTransactionsByBlockHash(blockhash string) ([]*entitiesweb3.Transaction, error) {
	iter := repo.DB.NewIterator(util.BytesPrefix([]byte("transaction_block_"+blockhash)), nil)
	defer iter.Release()

	var transactions []*entitiesweb3.Transaction
	for iter.Next() {
		var transaction entitiesweb3.Transaction
		err := json.Unmarshal(iter.Value(), &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repo *TransactionRepository) SaveTransaction(transaction *entitiesweb3.Transaction) error {
	transactionData, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("transaction_"+transaction.Hash), transactionData, nil)
	if err != nil {
		return err
	}

	err = repo.DB.Put([]byte("transaction_block_"+transaction.Hash), transactionData, nil)
	if err != nil {
		return err
	}

	return nil
}
