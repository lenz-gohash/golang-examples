// A TCP client that connects to the server at `address` and sends `message`.
package client

import (
	"context"
	"log"
	"net"
	"time"
)

// Dial connects to the server at the given `address` and sends the given `message`.
func Dial(address string, message string) {
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
