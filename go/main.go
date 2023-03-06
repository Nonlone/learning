package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

type EnumType int

const (
	A EnumType = iota + 1
	B
	C
	D
)

func printSliceSize(slice []int) {
	fmt.Printf("len=%d \n", len(slice))
}

func (enum EnumType) test2() {
	fmt.Println(enum)
}

func (enum *EnumType) test() {
	fmt.Println(*enum)
}

func test2(enum EnumType) {
	fmt.Println(enum)
}

func test3(enum *EnumType) {
	fmt.Println(*enum)
}

func main() {
	fmt.Println("hello world")

	var slice = make([]int, 3, 5)

	printSliceSize(slice)

	fmt.Println(A, B, C, D)

	b := EnumType(456)

	b.test()

	var a EnumType = 123
	var ap *EnumType

	ap = &a

	ap.test()

	var c EnumType = 789
	c.test2()

	testA := &TestA{}
	fmt.Println("testA", *testA)

	dsn := "root:123456@(127.0.0.1:3307)/test"

	gorm, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// var _ sql.DB = <-gorm.Statement.Context.Done();

	gorm.First(testA)

	fmt.Println("testa", *testA)

	xorm, err := xorm.NewEngine("mysql", dsn)
	xorm.ShowSQL(true)
	// err := engine.Sync(new(TestA))
	if err != nil {
		panic(err)
	}
	testA2 := &TestA{}
	xorm.ID(5).Get(testA2)
	fmt.Println("testA2", *testA2)

	sqlx, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	testA3 := TestA{}
	sqlx.Get(&testA3, "select * from testa where `key`='b'")

	fmt.Println("testA3", testA3)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("rand : %s", rand.Intn(1))

	var slice2 []int = make([]int, 0, 0)
	slice2 = append(slice2, 1)

	fmt.Printf("slice: %v \n", slice2)

	v := viper.New()
	v.SetConfigName("application-dev") // name of config file (without extension)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		// 配置错误
	}

	database := []databaseCluster{}
	v.UnmarshalKey("database",&database);

	jsonValue,_ := json.Marshal(database);

	fmt.Printf("database: %v \n", string(jsonValue))


	// fmt.Printf("viper.Get: %v \n",v.Get("database"));

	// DBInit()
	// RouterInit();

}

type databaseNode struct {
	Driver          string
	DatabaseSource string
	MaxIdleCount   int
	MaxOpen        int
	MaxLifetime    time.Duration
	MaxIdleTime    time.Duration
}

type databaseCluster struct { 
	Name string
	Master *databaseNode
	Slaves *[]databaseNode
}


