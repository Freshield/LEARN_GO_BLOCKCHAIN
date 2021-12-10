/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: check_type_convert.py
@Time: 2021-12-08 19:23
@Last_update: 2021-12-08 19:23
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package main

import "fmt"

//type Animal interface {
//	GetName() string
//}

type Cat struct {
	name string
}

func (c *Cat) GetName() string {
	return "I'm cat : " + c.name
}

type Dog struct {
	name string
}

func (d *Dog) GetName() string {
	return "I'm dog : " + d.name
}

func main() {
	cat := Cat{
		name: "hello kitty",
	}

	//animal := Animal(&cat)
	//fmt.Println(animal.GetName())

	dog1 := Dog(cat)
	fmt.Println(dog1.GetName())
}