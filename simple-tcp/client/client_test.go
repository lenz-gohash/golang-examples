package client

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
)

// FIXME
// Dial connects to the server at the given `address` and sends the given `message`.
func TestDial(t *testing.T) {
	var address = "localhost:12345"
	var message = "Hello, World!"

	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}

}
