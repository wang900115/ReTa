package entitiesweb3

type Block struct {
	Hash         string        // Hash of the block
	Index        int           // height of the block
	PreviousHash string        // Hash of the previous block
	Timestamp    int64         // Timestamp of the block
	Transactions []Transaction // List of transactions
	Miner        string        // Miner's address
	Difficulty   int
	Nonce        int    // Random number to prevent double spending
	MerkleRoot   string // Merkle root of the block
}
