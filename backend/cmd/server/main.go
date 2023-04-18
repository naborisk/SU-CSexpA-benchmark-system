package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/config"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/interfaces/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// config from env
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// mysql client
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	// build grpc server
	grpcServer := grpc.NewServer(grpc.WithDB(db))

	// launch grpc server
	listener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
