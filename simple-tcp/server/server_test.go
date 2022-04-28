package server

import (
	"net"
	"testing"
	"time"
)

// FIXME
func TestServe(t *testing.T) {

	const address = "localhost:12345"

	// Start the server in a goroutine. This allows for a non-blocking call to
	// `Serve()`.
	go Serve(address)

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", address)

	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		t.Fatal(err)
	} else {
		t.Log("Sent message")
	}

	// TODO: Check that the server received the correct message.

	conn.Close()
}

func ExampleServe() {
	const address = ":12345" // all interfaces on port 12345

	Serve(address)
}
