package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")

	rows, err := db.Query("SELECT id, name FROM users ")
	if err != nil {
		log.Fatal("error")
		log.Fatal(err)

	}
	defer rows.Close()

	for rows.Next() {
		 var name string
         var id int
        if err := rows.Scan(&id,&name); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("name:%s ,id:is %d\n", name, id)
	}

}
