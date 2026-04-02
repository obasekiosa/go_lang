package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	PORT = 4000
	ADDR = "127.0.0.1"
)

var dataGen = genData()

func main() {
	fullAddr := ADDR + ":" + strconv.Itoa(PORT)
	fmt.Println(fullAddr)
	var wg sync.WaitGroup
	wg.Go(func() {
		startServer(fullAddr)
	})

	time.After(3 * time.Second)

	wg.Go(func() {
		startClient(fullAddr)
	})

	wg.Wait()

}

func startServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	handleErr(err)
	defer ln.Close()
	buf := make([]byte, 1024)

	fmt.Println("Accepting Requests")
	conn, err := ln.Accept()
	setDeadLines(conn)
	handleErr(err)
	fmt.Println("Accepeted request from", conn.RemoteAddr())
	for {

		setDeadLines(conn)
		n, err := conn.Read(buf)
		handleErr(err)
		var data int
		if n > 0 {
			data = int(buf[0])
		}
		fmt.Println("Server Received:", data)

		outData := []byte{byte(dataGen())}
		setDeadLines(conn)
		n, err = conn.Write(outData)
		handleErr(err)
		fmt.Println("Server Sent:", outData)
	}
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	handleErr(err)
	setDeadLines(conn)

	buf := make([]byte, 1024)

	for {
		data := []byte{byte(dataGen())}
		setDeadLines(conn)
		conn.Write(data)
		fmt.Println("Client Sent:", data)
		setDeadLines(conn)
		n, err := conn.Read(buf)
		handleErr(err)
		if n > 0 {
			fmt.Println("Client Received:", int(buf[0]))
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func genData() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func setDeadLines(conn net.Conn) {
	_ = conn.SetDeadline(time.Now().Add(2 * time.Minute))
}
