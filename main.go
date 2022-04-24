package main

import (
	"database/sql"
	"log"

	apihandler "github.com/KhanSufiyanMirza/mini-aspire-API/apiHandler"
	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load Config : ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connot connect to database ", err)
	}
	store := db.NewStore(conn)
	server, err := apihandler.NewServer(config, store)
	if err != nil {
		log.Fatal("Connot Create  Server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Connot Connect to Server:", err)
	}

}
