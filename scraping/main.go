package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Error    error
	Response *http.Response
}

func newDB() *sql.DB {
	conn, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

const insertSql string = `insert into books values(name, publisher, author, date_of_issue, price, create_user_id, create_date) 
(?, ?, ?, ?, ?, 1, now())`

func main() {
	urls := []string{"", "", "", "", "", "", "", ""}

	db := newDB()
	var eg errgroup.Group
	eg.Go(func() error {

		_, err := db.Exec(insertSql)
		return err
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
