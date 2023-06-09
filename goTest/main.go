package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"os"
	"github.com/jasonming/go-fuckerr/try"

	"github.com/jmoiron/sqlx"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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
	v.SetConfigName("application") // name of config file (without extension)
	v.SetConfigType("properties")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		// 配置错误
	}

	database := []databaseCluster{}
	v.UnmarshalKey("database", &database)

	jsonValue, _ := json.Marshal(database)

	fmt.Printf("database: %v \n", string(jsonValue))

	springDataBase := v.Get("spring.shardingsphere.datasource")
	jsonValue, _ = json.Marshal(springDataBase)

	fmt.Printf("springDatabase: %v \n", string(jsonValue))

	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("develop"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"nacos-test-fs.inshopline.com",
			6801,
		),
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "application",
		Group:  "admin-app-service"})


	fmt.Printf("config content: %v \n",content)
	v = viper.New()
	v.SetConfigType("YAML")
	v.ReadConfig(strings.NewReader(content))
	
	databaseContent,_ := json.Marshal(v.Get("database"))
	fmt.Printf("viper.Get: %v \n",string(databaseContent));

	// DBInit()
	// RouterInit();


	fmt.Println()
	fmt.Println(try.Must(os.Getwd()))
	fmt.Println(try.Must(os.Executable()))

}

type databaseNode struct {
	Driver         string
	DatabaseSource string
	MaxIdleCount   int
	MaxOpen        int
	MaxLifetime    time.Duration
	MaxIdleTime    time.Duration
}

type databaseCluster struct {
	Name   string
	Master *databaseNode
	Slaves *[]databaseNode
}
