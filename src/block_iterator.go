package main

import (
	"github.com/boltdb/bolt"
)

type BlockIterator struct {
	CurrentHash []byte
	db          *bolt.DB
}

func (bi *BlockIterator) Next() *Block {
	var block *Block
	bi.db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("test1"))
		val := bk.Get([]byte("key"))
		block = Deserialize(val)
		return nil
	})
	return block
}
