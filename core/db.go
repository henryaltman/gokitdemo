package core

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	Db *sql.DB
	Mu sync.Mutex
}

var (
	db  *sql.DB
	app = &application{}
)

func init() {
	//fmt.Println("init")
	app.Db = newDB()
	//fmt.Println("db",app.Db)
}

func newDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/gokitdemo")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

func Instance() *application {

	if app.Db != nil {
		//fmt.Println("db instance is not null")
		if err := app.Db.Ping(); err != nil {
			app.Mu.Lock()
			app.Db = newDB()
			app.Mu.Unlock()
		}
		return app
	}
	app.Mu.Lock()
	app.Db = newDB()
	app.Mu.Unlock()
	return app
}
