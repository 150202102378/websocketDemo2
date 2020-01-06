package conf

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	conf           Config
	configFilePath string
	psqlDB         *gorm.DB
	psqlDBLock     sync.Mutex
)

type Config struct {
	ListenPort1 string
	ListenPort2 string
	ReactPort1  string
	ReactPort2  string
	ServerIP    string
	Psql        PsqlStruct
}
type PsqlStruct struct {
	Host, Port, DB, User, Passwd string
}

func init() {
	//initEnv()
	if configFilePath == "" {
		configFilePath = "./conf/config.toml"
	}
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		fmt.Println(err)
	}
}

func GetPsqlDB() *gorm.DB {
	if psqlDB != nil {
		return psqlDB
	}
	psqlDBLock.Lock()
	defer psqlDBLock.Unlock()
	if psqlDB != nil {
		return psqlDB
	}
	db := conf.Psql.DB
	host := conf.Psql.Host
	port := conf.Psql.Port
	user := conf.Psql.User
	passwd := conf.Psql.Passwd
	dbsql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, passwd, db,
	)
	var err error
	psqlDB, err = gorm.Open("postgres", dbsql)
	if err != nil {
		log.Fatal("init psql err : ", err)
		os.Exit(3)
	}
	return psqlDB
}

func initEnv() {
	flag.StringVar(&configFilePath, "c", "", "Init ConfigFilePath In Shell")
	flag.Parse()
}

func GetConfig() Config {
	return conf
}
