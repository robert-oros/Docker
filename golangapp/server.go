package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "html/template"
    _ "github.com/go-sql-driver/mysql"
)

func printRows(db *sql.DB) {
	var id, nume, prenume string

	row, _ := db.Query("SELECT id, nume, prenume FROM demisol")
	for row.Next() {
		row.Scan(&id, &nume, &prenume)
		log.Println("Person: ", id, " ", nume, " ", prenume)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "ADD")

	if r.Method == http.MethodGet {
		nume := r.URL.Query().Get("nume")
		prenume := r.URL.Query().Get("prenume")
		if err := addData(nume, prenume); err != nil {
			// err = fmt.Errorf("add data: %w", err)
			log.Printf("%s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func addData(nume string, prenume string) error {
    db, err := sql.Open("mysql", "tst:root@tcp(db:3306)/db")
	if err != nil {
		return fmt.Errorf("open mysql: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("ping: %w", err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS demisol (id INTEGER PRIMARY KEY, nume TEXT, prenume TEXT)")
	if err != nil {
		return fmt.Errorf("db prepare create: %w", err)
	}
	if _, err = statement.Exec(); err != nil {
		return fmt.Errorf("db exec: %w", err)
	}

	statement, err = db.Prepare("INSERT INTO demisol (nume, prenume) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("db prepare insert: %w", err)
	}

	if _, err = statement.Exec(nume, prenume); err != nil {
		return fmt.Errorf("db exec: %w", err)
	}
	printRows(db)
	return nil
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE")

	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		db, _ := sql.Open("mysql", "tst:root@tcp(db:3306)/db")
		statement, _ := db.Prepare("DELETE FROM demisol WHERE ID = ?")
		statement.Exec(id)

		printRows(db)
	}
}

func main() {
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/add", addHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
