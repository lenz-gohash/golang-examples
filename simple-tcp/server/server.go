// A tcp server that listens on a specified address and port.
package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// Start a server that listens on the specified address and port.
func Serve(address string) {
	// Listen on TCP `port` on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received connection from %s", conn.RemoteAddr())

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {

			for {

				// Read string until newline.
				line, err := bufio.NewReader(c).ReadString('\n')

				// Check for errors
				if err != nil {
					if errors.Is(err, io.EOF) { // prefered way by GoLang doc
						// Received EOF
						fmt.Println("Connection closed cleanly by peer")

					} else {
						// Check for other errors
						log.Fatal("Error reading: ", err.Error())
					}

					// Shut down the connection.
					c.Close()

					// Received Error or EOF
					break
				}

				// Log the received message and size
				// log.Printf("Received %d bytes: %s", len(bytes), string(bytes[:]))
				log.Printf("Received %d bytes: %s", len(line), line)
			}

		}(conn)
	}
}
