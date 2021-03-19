package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"path/filepath"
)

func NewDB() (db *sql.DB, err error) {
	type database struct {
		DriverName     string
		DataSourceName string
	}

	type Config struct {
		Database database
	}

	filePath, err := filepath.Abs("./configs/db.toml")
	if err != nil {
		return nil, err
	}
	conf := Config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("File contents: %s", content)
	toml.Unmarshal(content, &conf)

	log.Printf("%s", conf.Database.DriverName)
	log.Printf("%s", conf.Database.DataSourceName)

	db, err = sql.Open(conf.Database.DriverName, conf.Database.DataSourceName)
	return
}
