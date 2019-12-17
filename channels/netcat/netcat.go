package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to port 8000")

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Stopping Listening")
	}()
	toNetwork(conn, os.Stdin)
	conn.Close()
	log.Println("Program Terminated")
}

func toNetwork(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Println("Stopping Sending")
	}
}
