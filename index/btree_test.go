package index

import (
	"bottledb/data"
	"testing"
	"github.com/stretchr/testify/assert"
)


import TestBTree_Put(t *testing.T) {
	bt := NewBTree()

	st_put := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, st_put)

}