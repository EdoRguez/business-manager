package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/client-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/client-svc/pkg/db/sqlc"
	"github.com/EdoRguez/business-manager/client-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/client-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/client-svc/pkg/services"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	conn, err := sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	storage := db.NewStorage(conn)

	fmt.Println("Client Service ON: ", c.Port)

	// s := services.NewClientService(storage)
	s := services.ClientService{
		Repo: repo.NewClientRepo(storage),
	}

	grpcServer := grpc.NewServer()

	pb.RegisterClientServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
