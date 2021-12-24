/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: transaction.py
@Time: 2021-12-21 19:20
@Last_update: 2021-12-21 19:20
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

const subsidy = 10

type Transaction struct {
	ID []byte
	Vin []TXInput
	Vout []TXOutput
}

type TXInput struct {
	TXid []byte
	Vout int
	ScriptSig string
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

type TXOutput struct {
	Value int
	ScriptPubKey string
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func (tx *Transaction) SetId() {
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

func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].TXid) == 0 && tx.Vin[0].Vout == -1
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetId()

	return &tx
}

func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	// 如果累加完了依然小于额度则代表地址并没有这么多人的余额
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	// 遍历所有的输出，来转换为此交易的输入
	for txid, outs := range validOutputs {
		// 把txid变回16进制
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// 建立此交易的输出
	outputs = append(outputs, TXOutput{amount, to})
	// 如果输入大于输出则把找零返回给from的地址
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetId()

	return &tx
}