package index

import (
	"bottledb/data"
	"bytes"
	"github.com/google/btree"
)

// index abstract 
type IndexImpl interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (lhs *Item) Less(rhs btree.Item) bool {
	return bytes.Compare(lhs.key, rhs.(*Item).key) == -1
}