package index

import (
	"bottledb/data"
	"github.com/google/btree"
	"sync"
)

/**
	google/btree:
	thread unsafe while multi thread write
*/


// BTree encapsulation
type Btree struct {
	tree *btree.Btree
	lock *sync.RWMutex
}

func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),	// todo: config, btree subnode num
		lock: new(sync.RWMutex),
	}
}

func (bt *Btree) Put(key []byte, pos *data.LogRecordPos) bool {
	insert_item := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(insert_item)
	bt.lock.Unlock()
	return true
}


func (bt *Btree) Get(key []byte) *data.LogRecordPos {
	get_item_info := &Item{key: key}
	btree_item = bt.tree.Get(get_item_info)

	if btree_item == nil {
		return nil
	}

	return btree_item.(*Item).pos
}


func (bt *Btree) Delete(key []byte) del_result bool {
	del_item_info := &Item{key: key}

	bt.lock.Lock()
	oldItem = bt.tree.Delete(del_item_info)
	if oldItem == nil {
		del_result = false
	} else {
		del_result = true
	}
	bt.lock.Unlock()
}

