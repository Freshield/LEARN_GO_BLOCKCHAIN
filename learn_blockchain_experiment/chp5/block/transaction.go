/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: transaction.py
@Time: 2022-01-06 16:35
@Last_update: 2022-01-06 16:35
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
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// 默认的挖矿奖励
const subsidy = 10

type Transaction struct {
	// 交易的基本结构
	ID []byte
	Vin []TXInput
	Vout []TXOutput
}

func (tx Transaction) IsCoinbase() bool {
	// 判断是否为挖出来的奖励
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func (tx *Transaction) SetID() {
	// 得到此交易的hash并设为id
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

type TXInput struct {
	// 交易输入
	Txid []byte
	Vout int
	ScriptSig string
}

type TXOutput struct {
	// 交易输出
	Value int
	ScriptPubKey string
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	// 判断此交易输入是否可以解锁输出
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	// 判断此交易输出是否可以被解锁
	return out.ScriptPubKey == unlockingData
}

func NewCoinbaseTX(to, data string) *Transaction {
	// 新的挖矿交易
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	
	return &tx
}

func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	// 新的正常utxo交易
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}