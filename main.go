package simpleBlockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {

	var newBlock Block

	currentTime := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = currentTime.String()
	newBlock.BPM = BPM
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
