package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "10.40.1.12"
	port     = 5432
	user     = "iposselect"
	password = "'solo consultas'"
	dbname   = "sucursal"
)

func main() {
	fmt.Println("Connecting to database...")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	fmt.Println("Listing intellipos.terminal")
	// SELECT * para este ejemplo falla, hay que indicar las columnas a consultar
	rows, err := db.Query("SELECT id, nombre FROM intellipos.terminal")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var nombre string
		err = rows.Scan(&id, &nombre)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, nombre)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
