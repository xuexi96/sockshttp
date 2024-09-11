package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"sockshttp/consult"
	"sync"
)

func main() {
	//openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
	cert, err := tls.LoadX509KeyPair("./key/server.crt", "./key/server.key")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listen, err := tls.Listen("tcp", "0.0.0.0:8081", config)
	if err != nil {
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go Connectedhost(conn)
	}
}

func Connectedhost(conn net.Conn) {
	defer conn.Close()
	addr := consult.AnalyticConnection(conn)
	var status consult.Status
	if addr == (consult.Address{}) {
		return
	}
	dial, err := net.Dial(addr.Remoteprotocoltype, addr.RemoteAddr)
	if err != nil {
		consult.Response(conn, status)
		return
	}
	status.LocalAddr = dial.LocalAddr().String()
	status.Status = true
	consult.Response(conn, status)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(conn.LocalAddr().String() + "----------------------->" + dial.RemoteAddr().String())
		_, err := io.Copy(dial, conn)
		if err != nil {
			fmt.Println("写数据到", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = io.Copy(conn, dial)
		if err != nil {
			fmt.Println("读取数据", err)
			return
		}
	}()
	wg.Wait()

}
