package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// COMMANDS TO TESTE RESULT
	// $ go build gopl.io/ch8/netcat1
	// $ ./netcat1

	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_tcp2()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

func _tcp2() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
