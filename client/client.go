package client

import (
	"context"
	"log"
	"net"
	"time"
)

func Dial(address string, message string) {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// conn, err := d.DialContext(ctx, "tcp", "localhost:12345")
	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// if _, err := conn.Write([]byte("Hello, World!")); err != nil {
	if _, err := conn.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}
}
