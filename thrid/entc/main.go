package main

import (
	"context"
	stdSql "database/sql"
	"fmt"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"

	"github.com/buffge/gobyexample/thrid/entc/ent"
)

type Conf struct {
	AppName string `default:"ent demo"`
	Db      struct {
		User      string `default:"root"`
		Pwd       string `required:"true"`
		Ip        string
		Db        string
		Port      uint   `default:"3306"`
		ParseTime string `default:"True"`
		Loc       string `default:"Local"`
		Charset   string `default:"utf8mb4"`
	}
}

func ClientFromDB(db *stdSql.DB) *ent.Client {
	drv := sql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv))
}
func main() {
	conf := Conf{}
	configor.Load(&conf, "conf.json")
	db, err := stdSql.Open("mysql", fmt.Sprintf("%s:%s@tcp("+
		"%s:%d)/%s?parseTime=%s&charset=%s&loc=%s", conf.Db.User, conf.Db.Pwd, conf.Db.Ip,
		conf.Db.Port, conf.Db.Db, conf.Db.ParseTime, conf.Db.Charset, conf.Db.Loc))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	var client *ent.Client
	for i := 1; i < 1000; i++ {
		client = ClientFromDB(db)
		go func(client *ent.Client, i int) {
			for {
				<-time.After(1 * time.Second)
				_, err := client.Article.Get(context.Background(), i)
				if err == nil {
					log.Printf("%d success\n", i)
				}
			}
		}(client, i)
	}

	<-time.After(time.Hour)
	//str, _ := json.Marshal(art.ID)
	//fmt.Println(string(str))

}
