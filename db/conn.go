package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = 1234
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	//Construindo a string de conexão usando as constantes que definimos
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	//Abrindo a conexão com o postgres
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//Fazendo um ping para ver se a conexão abriu com sucesso
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	//Se sim, retornamos o db que é a nossa conexão com o banco
	return db, nil
}