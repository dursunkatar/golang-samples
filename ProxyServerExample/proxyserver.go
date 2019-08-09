package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buf       = make([]byte, 1024)
		r         = bufio.NewReader(conn)
		w         = bufio.NewWriter(conn)
		dataBytes []byte
	)

ILOOP:
	for {
		n, err := r.Read(buf)
		dataBytes = buf[:n]
		data := string(dataBytes)

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:\n", data)
			break ILOOP

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	response := SocketClient("127.0.0.1", 8080, dataBytes)
	w.Write(response)
	w.Flush()

}

func SocketClient(ip string, port int, data []byte) []byte {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}
	conn.Write(append(data, []byte("\r\n\r\n")...))
	result, _ := ioutil.ReadAll(conn)
	return result
}

func main() {
	port := 80
	SocketServer(port)
}
