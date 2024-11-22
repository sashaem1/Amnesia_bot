package telegram

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database = &sqlx.DB{}

type User struct {
	User_id  int64
	Role     string
	Username string
	Chat_id  int64
}

func DBConnection(host, port, dbname, user, password string) {
	var err any
	database, err = sqlx.Connect(
		"postgres", fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}
	err = database.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func DBCloseConnection() {
	database.Close()
}
