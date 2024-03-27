package model

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Nonce        int           `json:"nonce"`
	Hash         string        `json:"hash"`
	PrevHash     string        `json:"prevHash"`
}
