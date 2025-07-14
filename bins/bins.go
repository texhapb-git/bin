package bins

import (
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

// NewBin создает новый экземпляр Bin
func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

// NewBinList создает новый экземпляр BinList
func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0),
	}
}

// AddBin добавляет Bin в список
func (bl *BinList) AddBin(bin *Bin) {
	bl.Bins = append(bl.Bins, *bin)
}