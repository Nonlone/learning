package main

import (
	"fmt"

	"xorm.io/xorm"
)


var engine *xorm.Engine


type TestA struct {
	Id int64 
	Key string 
	Value string 
	
}

func (TestA) TableName() string{
	return "testa"
}


func DBInit(){
	fmt.Println("start init db...")
	engine,_=xorm.NewEngine("mysql","root:123456@127.0.0.1:3307/testa")
	// engine.ShowSQL(true)

}