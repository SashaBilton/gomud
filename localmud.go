package main

import (
	"bufio"
	"fmt"
	"gomud/entity"
	"gomud/interaction"
	"gomud/space"
	"os"
)

func main() {

	running := true

	reader := bufio.NewReader(os.Stdin)

	player := setupSimpleWorld()

	for running {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		cmdTokens := interaction.Tokenise(command)
		result := interaction.Do(cmdTokens, &player)
		fmt.Print(result)
		if cmdTokens[0] == "quit" {
			running = false
		}
	}
}

func setupSimpleWorld() entity.Player {
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

	player := entity.Player{Name: "Tester", Location: &start}
	return player
}
