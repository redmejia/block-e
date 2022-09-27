package blocke

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Blocke struct {
	Index    int
	Time     string
	Bpm      int
	Hash     string
	PrevHash string
}

type Blockes struct {
	B []Blocke
}

func NewBlocke() *Blocke {
	time := time.Now()
	return &Blocke{Index: 0, Time: time.String(), Bpm: 0, Hash: "", PrevHash: ""}
}

func NewBlockes() *Blockes {
	return &Blockes{B: []Blocke{}}
}

func (bs *Blockes) AddBlock(newBlock ...Blocke) {
	bs.B = append(bs.B, newBlock...)
}

// CalculateBlocke
func CalculateBlocke(b Blocke) string {

	bIndx := strconv.Itoa(b.Index)
	bBpm := strconv.Itoa(b.Bpm)
	recordHash := bIndx + b.Time + bBpm + b.PrevHash

	h := sha256.New()
	h.Write([]byte(recordHash))
	hash := h.Sum(nil)

	return hex.EncodeToString(hash)

}

// replaceChain
func (bs *Blockes) replaceChain(newBlock []Blocke) {
	if len(newBlock) > len(bs.B) {
		bs.B = newBlock
	}
}

// checkBlock
func checkBlock(newBlock *Blocke, oldBlock *Blocke) bool {
	if newBlock.Index+1 != oldBlock.Index {
		return false
	}

	if newBlock.Hash != oldBlock.Hash {
		return false
	}

	if CalculateBlocke(*newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// GenerateNewBlock
func (b *Blocke) GenerateNewBlock(old *Blocke, Bpm int) *Blocke {
	var newBlock Blocke

	t := time.Now()

	newBlock.Index = old.Index + 1
	newBlock.Time = t.String()
	newBlock.Bpm = Bpm
	newBlock.PrevHash = old.Hash
	newBlock.Hash = CalculateBlocke(newBlock)

	return &newBlock
}
