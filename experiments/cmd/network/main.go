package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	PORT          = 4000
	ADDR          = "127.0.0.1"
	READ_TIMEOUT  = 5
	WRITE_TIMEOUT = 5
	CLIENT_PREFIX = "CLIENT:"
	SERVER_PREFIX = "SERVER:"
)

var dataGen = genData()

func main() {
	fullAddr := ADDR + ":" + strconv.Itoa(PORT)
	fmt.Println(fullAddr)
	var wg sync.WaitGroup
	wg.Go(func() {
		startServer(fullAddr)
	})

	// some time for server to come online before client tries connecting
	<-time.After(3 * time.Second)

	wg.Go(func() {
		startClient(fullAddr)
	})

	wg.Wait()

}

// startServer and startClient operate async and so
// do not need to return values
// they do so and comunicate errors with cordinating primitives like channels
// failues should be handled by them and state communicated to main/coordinating process
// which now decides the enitre state of the program
// if startServer fails it should tell the starting process which now decides if to also end the client or
// try restarting the server

func startServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(SERVER_PREFIX, fmt.Errorf("Failed to start Server: %w", err))
		return
	}

	defer ln.Close()

	// logs
	fmt.Println(SERVER_PREFIX, "Server started on address: ", addr)
	fmt.Println(SERVER_PREFIX, "Accepting Requests")

	// Todo: set timer to trigger once connection count is zero
	// if it stays zero for lonkger that shutdown time without new connectin
	// then shutdown the server

	// to do this, we need to keep track of the connection count
	// anytime it hits zero (on creation and on decrement) start a timer
	// any time an increment happens clear the timer
	// on timer done, close server
	// keep in mind increments must first clear timer before setting up a new connection
	// to avoid the timer triggering anything, infact there are no guarantees that the clear
	// and setup operations in their own rights are atomic, so even if unlikely to have a cleared
	// timer pause execution and then the connection routine continues and then the clear operation
	// finishes but then the connection routine is still ongoing. it is best to coordinate all this
	// with the use of a mutex. A cleared timer must complete and release the mutex before any operation
	// that tires to start a connection

	for {
		conn, err := ln.Accept()
		if err != nil {
			// log
			fmt.Println(SERVER_PREFIX, fmt.Errorf("Could not accept connection request: %w", err))
			// keep track of failure rate and terminate based on treshold.
			continue
		} else {
			// log
			fmt.Println(SERVER_PREFIX, "Accepeted request from", conn.RemoteAddr())

			go handleServerConn(conn)
		}

	}
}

func handleServerConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	connPrefix := SERVER_PREFIX + " " + conn.LocalAddr().String()
	for {
		readData, err := handleRead(conn, r)
		if err == io.EOF {
			fmt.Println(connPrefix, fmt.Errorf("Got no data on attempted read. Terminating connection: %w", err))
			return
		} else if err != nil {
			// log, no data gotten this time
			// this isn't a state defined protocol so we can't decide if this is normal or not
			// for now we simply log it and try again
			fmt.Println(connPrefix, fmt.Errorf("Got no data on attempted read: %w", err))
		} else {
			fmt.Println(connPrefix, "Server Received:", readData)
		}

		setWriteDeadLine(conn)
		toSend := uint8(dataGen())
		// converting 64/32/... sized int to 1 byte (8 bits) is intentional
		_, err = conn.Write([]byte{byte(toSend)})

		// logs
		if err != nil {
			fmt.Println(connPrefix, fmt.Errorf("Failed to send %d :%w", toSend, err))
		} else {
			fmt.Println(connPrefix, "Server Sent:", toSend)
		}
	}
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(CLIENT_PREFIX, fmt.Errorf("Failed to start Client: %w", err))
		return
	}
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		data := uint8(dataGen())
		setWriteDeadLine(conn)
		_, err := conn.Write([]byte{byte(data)})

		if err != nil {
			fmt.Println(CLIENT_PREFIX, fmt.Errorf("Could not write %d: %w", data, err))
		} else {
			fmt.Println(CLIENT_PREFIX, "Client Sent:", data)
		}

		// how does readDeadline work with this read does the read block until it sees a byte?
		// or does it try and then leave if it doesn't?
		// oh conn reads/writes block until deadlines
		setReadDeadLine(conn)
		readData, err := r.ReadByte()
		// connections only end by network errors or server close or client close
		if err == io.EOF {
			fmt.Println(CLIENT_PREFIX, fmt.Errorf("Got no data on attempted read. Terminating connection: %w", err))
			return
		} else if err != nil {
			fmt.Println(CLIENT_PREFIX, fmt.Errorf("No data was gotten: %w", err))
		} else {
			fmt.Println(CLIENT_PREFIX, "Client Received:", readData)
		}
	}
}

func handleRead(conn net.Conn, reader *bufio.Reader) (int, error) {
	setReadDeadLine(conn)
	data, err := reader.ReadByte()
	if err != nil {
		return -1, err
	}
	return int(data), nil
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

func setWriteDeadLine(conn net.Conn) {
	err := conn.SetWriteDeadline(time.Now().Add(WRITE_TIMEOUT * time.Second))
	if err != nil {
		fmt.Println(fmt.Errorf("Could not set Write Deadline for conn %s <-> %s : %w", conn.LocalAddr(), conn.RemoteAddr(), err))
	}
}

func setReadDeadLine(conn net.Conn) {
	err := conn.SetReadDeadline(time.Now().Add(READ_TIMEOUT * time.Second))
	if err != nil {
		fmt.Println(fmt.Errorf("Could not set Read Deadline for conn %s <-> %s : %w", conn.LocalAddr(), conn.RemoteAddr(), err))
	}
}
