package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/kutyrov/my_server_go/internal/app/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8080", "bind port")
}

func main() {
	flag.Parse()
	// fmt.Println(port)
	bind_port := "8080"
	if _, err := strconv.Atoi(port); err != nil {
		log.Println("порт невалиден, выставлен дефолтный 8080")
	} else {
		bind_port = port
	}

	// fmt.Println(bind_port)

	my_server := server.New(bind_port)

	if err := my_server.Start(); err != nil {
		log.Fatal(err)
	}

}
