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

	<-time.After(3 * time.Second)

	wg.Go(func() {
		startClient(fullAddr)
	})

	wg.Wait()

}

func startServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	handleErr(err)
	defer ln.Close()

	fmt.Println("Accepting Requests")

	for {
		conn, err := ln.Accept()
		handleErr(err)
		go handleServerConn(conn)
		fmt.Println("Accepeted request from", conn.RemoteAddr())
	}
}

func handleServerConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4)
	for {
		readData := handleRead(conn, buf)
		fmt.Println("Server Received:", readData)

		<-time.After(3 * time.Second)

		outData := []byte{byte(dataGen())}
		setDeadLines(conn)
		_, err := conn.Write(outData)
		handleErr(err)
		fmt.Println("Server Sent:", outData)
	}
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	handleErr(err)
	setDeadLines(conn)

	buf := make([]byte, 4)

	for {
		data := []byte{byte(dataGen())}
		setDeadLines(conn)
		_, err = conn.Write(data)
		handleErr(err)
		fmt.Println("Client Sent:", data)

		readData := handleRead(conn, buf)
		fmt.Println("Client Received:", readData)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handleRead(conn net.Conn, buf []byte) []int {
	setDeadLines(conn)
	<-time.After(3 * time.Second)
	n, err := conn.Read(buf)
	handleErr(err)
	out := make([]int, n)
	for i := range n {
		out[i] = int(buf[i])
	}
	return out
}

func genData() func() int {
	i := 0
	var lock sync.Mutex

	return func() int {
		lock.Lock()
		defer lock.Unlock()
		i += 1
		return i
	}
}

func setDeadLines(conn net.Conn) {
	_ = conn.SetDeadline(time.Now().Add(2 * time.Minute))
}
