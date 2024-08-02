package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/KelpGF/Go-gRPC/internal/database"
	"github.com/KelpGF/Go-gRPC/internal/pb"
	"github.com/KelpGF/Go-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id TEXT, name TEXT, description TEXT)")
	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	log.Println("Server is running at port 50051")

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
