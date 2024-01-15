package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME = "todo.db"
	DB_PATH = "./" + DB_NAME
)

func main() {
	createDBIfNotExist()
	db := openDB()

	createTableIfNotExist(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run() // listen and serve on
}

func createDBIfNotExist() {
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		file, err := os.Create(DB_PATH)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", DB_PATH)

	if err != nil {
		panic(err)
	}

	return db
}

func createTableIfNotExist(db *sql.DB) {
	sqlStmt := `
	create table if not exists foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err := db.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}
}
