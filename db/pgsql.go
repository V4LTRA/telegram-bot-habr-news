package db

import (
	"database/sql"
	"fmt"
	"news-bot/logger"
	"news-bot/telegram"
	"time"
	_ "github.com/lib/pq"
)

/*func mysqlConn(dbName string) *sql.DB {
	db, err := sql.Open("mysql", "root:Passwd@unix(/var/run/mysqld/mysqld.sock)/%s", dbName))
	if err != nil {
		panic(err.Error())
	}
	return db 
	
}
*/

func PostConn() *sql.DB {
Connect := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", Connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func selectHash(db *sql.DB, hash string) bool {
	var rowHash int
	db.QueryRow(`SELECT id FROM tg."News" WHERE link_hash = $1`, hash).Scan(&rowHash)
	if rowHash != 0 {
		return true
	}
	return false
}

func insertHash(db *sql.DB, url, page, text, hash string, times time.Time) {
	resu, err := db.Prepare(`INSERT INTO tg."News" (site, page_link, page_text, timestamp, link_hash) VALUES ($1, $2, $3, $4, $5)`)
	fmt.Println(err)
	if err != nil {
		logger.ForError(err)
	}
	fmt.Println("Новая новость с:")
	_, err = resu.Exec(url, page, text, times, hash)
	if err != nil {
		logger.ForError(err)
	}
}

func CheckSiteNewsBot(url, page, text, hash string) {
	
	database := PostConn()
	
	checkLink := selectHash(database, hash)
	
	times := time.Now()

	if checkLink == false {
		insertHash(database, url, page, text, hash, times)
		telegram.SendMessage(text)
	}
	defer database.Close()
}
