package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
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

const maxProcess int = 20

func main() {
	urls := []string{"", "", "", "", "", "", "", ""}

	urlChan := make(chan string, len(urls))

	//db := newDB()
	var eg errgroup.Group

	for i := 0; i < maxProcess; i++ {
		eg.Go(func() error {
			for url := range urlChan {
				resp, err := http.Get(url)
				if err != nil {
					return err
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				_ = resp.Body.Close()

				fmt.Println(string(body))
				//_, err = db.Exec(insertSql)
				//if err != nil {
				//	return err
				//}
				//return err
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
