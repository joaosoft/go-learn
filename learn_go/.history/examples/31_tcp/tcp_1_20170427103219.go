package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// COMMANDS TO TESTE RESULT
	// $ go build gopl.io/ch8/clock1
	// $ ./clock1 &
	// $ nc localhost 8000

	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_tcp()
	Main1()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

func _tcp() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
