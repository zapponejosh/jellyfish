package dbOps

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "jellyapp"
	password = "fish"
	dbname   = "jelly"
)

func init() {
	db = connectDB(dbname)

	defer db.Close()

}
func dsn(dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
}

func connectDB(dbname string) *sql.DB {

	db, err := sql.Open("postgres", dsn(dbname))
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	return db
}

func TestQuery() string {
	db := connectDB(dbname)
	defer db.Close()
	rows, err := db.Query("select * from test")
	if err != nil {
		panic(err.Error())
	}
	var testStr string
	for rows.Next() {
		err = rows.Scan(&testStr)
		if err != nil {
			panic(err.Error())
		}
	}
	return testStr
}
