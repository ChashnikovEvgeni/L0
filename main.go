package main

import (
	"L0/server"
	"L0/service"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	var (
		connStr   string
		clusterID string
		clientID  string
		port      string
	)

	//fmt.Println(connStr, clusterID, clientID, port)

	service := service.Service{}
	service.Set_config(connStr, clusterID, clientID)
	service.DB_connect()
	service.Make_start_cache()
	server := server.Httpserver{}
	server.Set_config(&port, &service)
	go service.Connect_listen()
	go server.Server_work()
	fmt.Scanln()

}
