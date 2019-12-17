package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Listening in port 8000")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, voice string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(voice))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", voice)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(voice))
}
