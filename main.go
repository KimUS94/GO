package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	db, err := sql.Open("mssql", "server=(local);user id=semmlcc;password=semmlcc;database=SEMMLCC_345_0727")
	//db, err := sql.Open("mssql", "host=127.0.0.1 port=1433 user=semmlcc password=semmlcc dbname=SEMMLCC_345_0727 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리 : QueryRow()
	var lname, fname string

	err = db.QueryRow("SELECT UPDATETIME, UPDATE_USER FROM JIG_BTM_STATE_STD").Scan(&lname, &fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fname, lname)

}

//http://golang.site/go/article/110-MS-SQL-Server-%EC%82%AC%EC%9A%A9
