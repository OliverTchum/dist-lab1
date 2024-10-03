package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	conn, _ := ln.Accept()
	conns <- conn
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	stdin := bufio.NewReader(client)
	for {
		msg, _ := stdin.ReadString('\n')
		newmessage := Message{clientid, msg}
		msgs <- newmessage

	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()
	ln, _ := net.Listen("tcp", *portPtr)
	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	clientID := 0
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
			go handleClient(conn, clientID, msgs)
			clients[clientID] = conn
			clientID++
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i, c := range clients {
				if i != msg.sender {
					fmt.Fprintln(c, msg.message)
				}

			}

		}
	}
}
