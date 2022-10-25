package main

import (
	"L0/server"
	"L0/service"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var usageStr = `

Options:
	-conn database connection datar example: user=postgres password=123 dbname=L0 sslmode=disable
	-c  The NATS Streaming cluster ID  example: cluster
	-id The NATS Streaming client ID to connect with example:id
	-p Port for server  example:  :8084

`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func main() {

	var (
		connStr   string
		clusterID string
		clientID  string
		port      string
	)
	flag.StringVar(&connStr, "conn", "user=postgres password=123 dbname=L0 sslmode=disable", "database connection datar ")
	flag.StringVar(&clusterID, "c", "cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "id", "The NATS Streaming client ID to connect with")
	flag.StringVar(&port, "p", ":8084", "Port for server")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	service_l := service.Service{}
	service_l.Set_config(connStr, clusterID, clientID)
	service_l.DB_connect()
	service_l.Make_start_cache()
	server := server.Httpserver{}
	server.Set_config(&port, &service_l)
	go service_l.Connect_listen()
	go server.Server_work()
	fmt.Scanln()

}
