package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func Encode(block *Block) []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func Decode(bs []byte) *Block {
	var block Block

	reader := bytes.NewReader(bs)
	decoder := gob.NewDecoder(reader)
	decoder.Decode(&block)

	return &block
}

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("test1"))
		if err != nil {
			log.Panic(err)
		}
		return b.Put([]byte("key"), []byte("deserve"))
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("test1"))
		val := b.Get([]byte("key"))
		fmt.Println(string(val))

		return nil
	})
}
