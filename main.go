package main

import (
	"github.com/ktruedat/proxy-server/utils"
	"io"
	"net"
)

var errLog = utils.NewErrLogger()

func main() {

	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		errLog.Fatalf("Failed to initialize listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			errLog.Fatalf("Unable to accept connection: %v", err)
		}
		go handleConn(conn)
	}

}

func handleConn(src net.Conn) {
	//Requests coming to google.com, for example
	dst, err := net.Dial("tcp", "www.google.com:443")
	if err != nil {
		errLog.Fatalf("Unable to connect to target server: %v", err)
	}

	defer func(dst net.Conn) {
		err := dst.Close()
		if err != nil {
			errLog.Fatalf("Unable to close conn: %v", err)
		}
	}(dst)

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			errLog.Fatal(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		errLog.Fatal(err)
	}
}
