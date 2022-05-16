package main

import (
	"database/sql"
	"log"
	"net"

	apihandler "github.com/KhanSufiyanMirza/mini-aspire-API/apiHandler"
	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/gapi"
	"github.com/KhanSufiyanMirza/mini-aspire-API/ma"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	// go runGinServer(config, store)

	runGrpcServer(config, store)

}
func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Connot Create  Server:", err)
	}

	grpcServer := grpc.NewServer()
	ma.RegisterMiniAspireServer(grpcServer, server)
	reflection.Register(grpcServer)
	listner, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Connot create listner:", err)
	}
	log.Printf("start gRPC  server at %s", listner.Addr().String())
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("Connot start gRPC Server :", err)
	}
}
func runGinServer(config utils.Config, store db.Store) {
	server, err := apihandler.NewServer(config, store)
	if err != nil {
		log.Fatal("Connot Create  Server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Connot Connect to Server:", err)
	}
}
