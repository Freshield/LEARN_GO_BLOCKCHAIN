/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: blockchain_iterator.py
@Time: 2021-12-17 17:44
@Last_update: 2021-12-17 17:44
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package block

import "github.com/boltdb/bolt"

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodeBlock)

		return nil
	})
	if err != nil {
		panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
