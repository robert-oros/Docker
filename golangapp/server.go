package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	DB *sql.DB
}

func newService(DB *sql.DB) *Service {
	return &Service{DB: DB}
}

func connDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "tst:root@tcp(db:3306)/db")
	if err != nil {
		return db, fmt.Errorf("open mysql: %w", err)
	}

	if err = db.Ping(); err != nil {
		return db, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}

func main() {
	db, err := connDB()
	if err != nil {
		fmt.Println("connDB %s", err)
	}

	s := newService(db)

	http.HandleFunc("/delete", s.deleteHandler)
	http.HandleFunc("/add", s.addHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (s *Service) addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		nume := r.URL.Query().Get("nume")
		prenume := r.URL.Query().Get("prenume")
		if err := addData(nume, prenume, s.DB); err != nil {
			log.Printf("%s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *Service) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		
		statement, err := s.DB.Prepare("DELETE FROM demisol WHERE ID = ?")
		if err != nil {
			fmt.Println("delete from %s", err)
		}

		statement.Exec(id)

		err = printRows(s.DB)
		if err != nil {
			fmt.Println("print rows %s", err)
		}
	}
}

func addData(nume, prenume string, DB *sql.DB) error {
	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS demisol (id INTEGER PRIMARY KEY, nume TEXT, prenume TEXT)")
	if err != nil {
		return fmt.Errorf("db prepare create: %w", err)
	}
	if _, err = statement.Exec(); err != nil {
		return fmt.Errorf("db exec: %w", err)
	}

	statement, err = DB.Prepare("INSERT INTO demisol (nume, prenume) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("db prepare insert: %w", err)
	}

	if _, err = statement.Exec(nume, prenume); err != nil {
		return fmt.Errorf("db exec: %w", err)
	}

	err = printRows(DB)
	if err != nil {
		fmt.Println("print rows %s", err)
	}

	return nil
}

func printRows(db *sql.DB) error {
	var id, nume, prenume string

	row, err := db.Query("SELECT id, nume, prenume FROM demisol")
	if err != nil {
		return fmt.Errorf("select frm demisol: %w", err)
	}

	for row.Next() {
		row.Scan(&id, &nume, &prenume)
		log.Println("Person: ", id, " ", nume, " ", prenume)
	}

	return nil
}
