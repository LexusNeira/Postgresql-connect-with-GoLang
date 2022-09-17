package main

// query Postgresql database version
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234qwer"
	dbname   = "postgres"
)

var (
	version string
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", version)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
