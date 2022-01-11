/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: utils.py
@Time: 2022-01-06 16:33
@Last_update: 2022-01-06 16:33
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
	"encoding/binary"
)

func IntToHex(num int64) []byte {
	// 把int值转为hex值
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}