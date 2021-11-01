package main

import (
	"fmt"
	"gomud/entity"
	"gomud/interaction"
	"gomud/space"
	"math/rand"
	"net"
	"os"
	"strconv"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	start := setupSimpleWorld()
	playerLocations := new(entity.PlayerMap)
	playerLocations.PL = make([]entity.PlayerLocation, 0, 100)
	fmt.Println(&playerLocations)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		fmt.Printf("%T\n", playerLocations)
		go handleRequest(conn, &start, playerLocations)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, start *space.Location, playerLocations *entity.PlayerMap) {

	running := true

	name := "Wanderer " + strconv.Itoa(rand.Intn(999))
	fmt.Println(name + " joined.")
	fmt.Println(&playerLocations)
	player := entity.Player{Name: name, Location: start}
	playerLocations.MovePlayer(start, &player)

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	for running {
		conn.Write([]byte("> "))
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		command := string(buf[:])
		cmdTokens := interaction.Tokenise(command)
		result := interaction.Do(cmdTokens, &player, playerLocations)
		conn.Write([]byte(result))
		if cmdTokens[0] == "quit" {
			running = false
		}
		if cmdTokens[0] == "sd" {
			fmt.Println("shutdown initiated.")
			os.Exit(1)
		}
		if cmdTokens[0] == "debug" {
			msg := "\nplayer-locations\n"
			conn.Write([]byte(msg))
			conn.Write([]byte(entity.List(playerLocations)))

			conn.Write([]byte(msg))
		}
	}

	conn.Close()
}

func setupSimpleWorld() space.Location {
	startDesc := "This is where you start. It's groovy in here. "

	start := space.Location{Desc: startDesc}

	endDesc := "This is where you end"
	end := space.Location{Desc: endDesc}
	exit := space.Exit{Name: "end", Desc: "An exit that lends to the end.", Location: &end}

	homeDesc := "You're home. Safe, warm and filled with all the best things"
	home := space.Location{Desc: homeDesc}
	homeExit := space.Exit{Name: "home", Desc: "A glowing tunnel that leads home.", Location: &home}
	startExit := space.Exit{Name: "start", Desc: "A glowing tunnel that leads to the start.", Location: &start}

	start.AddExit(exit)
	start.AddExit(homeExit)
	home.AddExit(startExit)

	return start

}
