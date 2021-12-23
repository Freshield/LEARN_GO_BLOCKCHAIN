/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: blockchain.py
@Time: 2021-12-21 16:02
@Last_update: 2021-12-21 16:02
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
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const (
	dbFile = "blockchain.db"
	blocksBucket = "blocks"
	genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
)

type Blockchain struct {
	tip []byte
	DB *bolt.DB
}

func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lashHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lashHash = b.Get([]byte("l"))

		return nil
	})

	newBlock := NewBlock(transactions, lashHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func NewBlockchain() *Blockchain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

func CreateBlockchain(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.DB}

	return bci
}

func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	// 得到此地址的所有未花费输出交易
	var unspentTxs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		// 得到这个block
		block := bci.Next()
		// 遍历block中的每一个交易
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			// 遍历交易中的所有TXOutput
			for outIdx, out := range tx.Vout {
				// 如果交易的id在已花费的UTXO字典中找到了
				if spentTXOs[txID] != nil {
					// 遍历交易的id在已花费UTXO中的列表
					for _, spentOut := range spentTXOs[txID] {
						// 如果找到了这笔交易的索引index已经在列表中了
						// 表明这笔交易的输出已经被使用过了则跳过此TXOutput
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				// 如果此TXOutput没被使用且可以被此地址解锁
				// 那么便是当前地址的未花费输出
				if out.CanBeUnlockedWith(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
			// 开始看除去CB的输入
			if tx.IsCoinbase() == false {
				// 遍历每笔交易的每个输入
				for _, in := range tx.Vin {
					// 如果来自本地址，则代表此输入对应的输出已被花费
					if in.CanUnlockOutputWith(address) {
						// 加到已花费的字典中
						inTxID := hex.EncodeToString(in.TXid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}
		// 如果反推到了创世块则退出
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTxs
}

func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	// 找到所有未花费的输出
	unspentOutputs := make(map[string][]int)
	// 找到所有未花费输出的交易
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0
Work:
	for _, tx := range unspentTXs {
		// 得到此交易的ID
		txID := hex.EncodeToString(tx.ID)
		// 遍历所有的输出
		for outIdx, out := range tx.Vout {
			// 如果输出可以被此地址解锁且累加的金额还比需要的金额小
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				// 继续累加金额并把相应的交易信息放到map中
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
				// 如果累加金额已经大于了所需要的金额则跳出
				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)
	// 遍历所有未花费的交易
	for _, tx := range unspentTransactions {
		// 遍历所有TXOutput，如果可以被address解锁，则代表为此地址的未花费输出
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}