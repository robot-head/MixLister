/* MixLister server.
 * 
 * This scrapes shoutcast metadata and discards the audio
 *
 */
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var serverPort *int64 = flag.Int64("l", 8000, "The port to listen for connections on.")

func main() {

	flag.Parse() // Scan the arguments list 

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	service := ":" + strconv.FormatInt(*serverPort, 10)
	tcpAddr, err := net.ResolveTCPAddr("ip4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		buf := make([]byte, 100)
		stop := false
		for !stop {
			n, err := conn.Read(buf[0:])
			if err != nil {
				stop = true
			}
			os.Stdout.Write(buf[0:n])
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
