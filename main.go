package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gcfg.v1"
)

var (
	datasource Repository
)

type Config struct {
	General struct {
		Port       string
		StorageDir string
	}
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}
}

func main() {

	// config
	var (
		mysqlLocation string
		config        Config
	)
	err := gcfg.ReadFileInto(&config, "config.gcfg")
	if err != nil {
		log.Printf("Error getting config information: %s", err.Error())
		panic(err)
	}

	// mySql information for login
	mysqlLocation = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database,
	) // "root@tcp(127.0.0.1:3306)/test"

	datasource = NewDBSource(mysqlLocation)
	// defer datasource.purldb.db.Close()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
