/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: base58.py
@Time: 2022-01-14 20:40
@Last_update: 2022-01-14 20:40
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
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// 把输入专为base58编码
func Base58Encode(input []byte) []byte {
	var result []byte

	// 把input转化为int
	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	// 如果x不是0
	for x.Cmp(zero) != 0 {
		// x对base取模，同时x变为除以base后的数
		x.DivMod(x, base, mod)
		// 对应的字符添加到结果中
		result = append(result, b58Alphabet[mod.Int64()])
	}

	// 把结果反过来，也就是变为正序
	ReverseBytes(result)
	// 遍历输入，如果从开头有0则补上，因为变为int的时候会忽略
	for b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

// 把base58转为输入
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	// 看有多少个0
	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	// 去除开头为0的部分
	payload := input[zeroBytes:]
	for _, b := range payload {
		// 得到b在字母表中的索引位置
		charIndex := bytes.IndexByte(b58Alphabet, b)
		// 取模的反向操作
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	// 把开头的0的部分补上
	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)

	return decoded
}