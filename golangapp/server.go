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
	fmt.Fprintf(w, "ADD")

	if r.Method == http.MethodGet {
		nume := r.URL.Query().Get("nume")
		prenume := r.URL.Query().Get("prenume")
		addData(nume, prenume)
	}
}

func addData(nume string, prenume string) {
    db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:4000)/db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS demisol (id INTEGER PRIMARY KEY, nume TEXT, prenume TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO demisol (nume, prenume) VALUES (?, ?)")

	statement.Exec(nume, prenume)
	printRows(db)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE")

	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:4000)/db")
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
