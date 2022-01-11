/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: blockchain.py
@Time: 2022-01-11 17:13
@Last_update: 2022-01-11 17:13
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

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// 链的基本结构
type Blockchain struct {
	tip []byte
	db *bolt.DB
}

// 链数据的迭代器
type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

// 挖区块
func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lashHash []byte
	// 得到last区块
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lashHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	// 挖新区块
	newBlock := NewBlock(transactions, lashHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

// 得到链的迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// 得到下一个区块，其实是上一个
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

// 找到未花费的交易
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	// 为花费的交易
	var unspentTXs []Transaction
	// 花费了的交易，key是交易的id，value为交易的索引
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		// 得到当前的区块
		block := bci.Next()

		// 遍历当前区块的交易
		for _, tx := range block.Transactions {
			// 得到当前区块的id
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			// 从out入手
			// 遍历交易的out部分
			for outIdx, out := range tx.Vout {
				// 如果当前的交易id已经在已花费字典中有值
				if spentTXOs[txID] != nil {
					// 遍历已花费字典，看当前的out的索引和已花费字典记录的索引一致
					// 表明此out已经被花费了，继续下一个out
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				// 如果当前交易id没有再已花费字典中
				// 或有交易id但是out的索引并不在其中
				// 且可以被当前地址解锁，则表明为当前地址的未花费交易
				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}
			// 如果当前交易不是挖块交易
			// 则遍历当前交易的所有in的部分
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					// 如果当前的in部分被当前地址解锁了
					// 则把当前交易的id和索引位置放到已花费字典中
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}
		// 如果已经到了创世区块则退出
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

// 找到所有的未花费输出
func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// 找到所有的可以花费的输出
func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}