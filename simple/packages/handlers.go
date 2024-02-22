package simple

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


const (
	host = "localhost"
	port = 5432
)

type Item struct {
	TaskNum int `json:"id"`
	Task string `json:"task"`
	Status bool `json:"status"`
}

func OpenConnection() (*sql.DB, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	user, ok := os.LookupEnv("USER")
    if !ok {
        log.Fatal("Error loading env variables")
    }
    password, ok := os.LookupEnv("PASSWORD")
    if !ok {
        log.Fatal("Error loading env variables")
    }
	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("Error loading env variables")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	email := GetEmail()
	addEmail := `insert into users (email) values($1) on conflict (email) do nothing;`
	_, err = db.Exec(addEmail, email)
	if err != nil {
		panic(err)
	}

	var userId string
	getUser := `select user_id from users where email = $1;`
	err = db.QueryRow(getUser, email).Scan(&userId)
	if err != nil {
		panic(err)
	}

	return db, userId
}

func GetEmail() string {
	return ""
}

var GetList = http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, userId := OpenConnection()

	rows, err := db.Query("select id, task, status from tasks join users on task.user_uuid = users.user_id where user_id = $1;", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}
	defer rows.Close()
	defer db.Close()

	items := make([]Item, 0)

	for rows.Next() {
		var item Item 
		err := rows.Scan(&item.TaskNum, &item.Task, &item.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err)
		}
		items = append(items, item)
	}
	
	itemBytes, _ := json.MarshalIndent(items, "", "\t")

	_, err = w.Write(itemBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	// _ = json.NewEncoder(w).Encode(items)
})

