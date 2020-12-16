package configs

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

type database struct {
	DriverName     string
	DataSourceName string
}

type Config struct {
	Database database
}

func TestToml(t *testing.T) {
	filePath, err := filepath.Abs("db.toml")
	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("File contents: %s", content)
	toml.Unmarshal(content, &conf)
	fmt.Printf("%#v\n", conf)
}

func TestToml1(t *testing.T) {

	type Postgres struct {
		User     string
		Password string
	}
	type Config struct {
		Postgres Postgres
	}

	doc := []byte(`
[Postgres]
User = "pelletier"
Password = "mypassword"`)

	config := Config{}
	toml.Unmarshal(doc, &config)
	fmt.Printf("user=%v", config)
}
