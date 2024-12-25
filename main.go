/* GetHeadInfo
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	// port number
	service := "10.0.0.74:1300"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		location, err := time.LoadLocation("Asia/Tokyo")
		checkError(err)
		japanTime := time.Now().In(location).String()
		conn.Write([]byte(fmt.Sprintf("%s\n", japanTime)))
		conn.Close()
	}
}
