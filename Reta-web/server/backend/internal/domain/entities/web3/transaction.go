package entitiesweb3

type Transaction struct {
	Sender    string // Sender's Nickname
	Receiver  string // Receiver's Nickname
	Amount    float64
	Fee       float64 // Transaction fee
	Hash      string
	Signature string
	Nonce     int // Random number to prevent double spending
	Status    string
}
