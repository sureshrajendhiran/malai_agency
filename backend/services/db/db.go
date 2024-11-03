package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

var DB *sql.DB

// MongDB its hold mongodb connection
var MongoDB *mgo.Database
var DataBaseName string

func Init() {
	// DataBaseName = "flow_pod"
	DataBaseName = "flow_pod1"
	ConnectionMake()
}

func ConnectionMake() {

	db, err := sql.Open("mysql", "root:alu&4321@tcp(localhost:3306)/"+DataBaseName)
	// db, err := sql.Open("mysql", "root:Suresh@1234@tcp(localhost:3306)/"+DataBaseName)
	fmt.Println(db, err)
	if err != nil {
		log.Println("Database Connection Error:", err)
	}
	// db.SetMaxOpenConns(5)
	DB = db
}

// MysqlConObj func used to get mysql Connection
func MysqlConObj() *sql.DB {
	return DB
}
