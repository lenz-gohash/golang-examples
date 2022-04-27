package server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
)

// FIXME
func TestServe(t *testing.T) {

	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}

		log.Printf("Received connection from %s", conn.RemoteAddr())

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.

		// Use a channel to hand off errors
		errs := make(chan error)
		go func(c net.Conn) {

			// Make a buffer to hold incoming message
			buffer := make([]byte, 1024)

			// Read the incoming connection into the buffer.
			bytes, err := c.Read(buffer) // read the message

			// Check for errors
			if err != nil {
				if errors.Is(err, io.EOF) { // prefered way by GoLang doc
					// Check for EOF
					fmt.Println("Finished reading from client: Received EOF")
				} else {
					// Check for other errors

					// t.Fatal("Error reading: ", err.Error())
					errs <- err
				}
			}

			// Log the received message and size
			log.Printf("Received %d bytes: %s", bytes, string(buffer[:bytes]))

			// Shut down the connection.
			c.Close()
		}(conn)

		// Check for errors
		err = <-errs
		if err != nil {
			t.Fatal("Error reading: ", err.Error())
		}
	}

}
