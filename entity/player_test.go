package entity

import (
	"gomud/space"
	"testing"
	"time"
)

//Move the player to a named exit based on that players current location
//This means Player, Location & Exit structures need to be created
func TestMoveToNamedExit(t *testing.T) {

	startDesc := "This is where you start"

	start := space.Location{Desc: startDesc}

	endDesc := "This is where you end"
	end := space.Location{Desc: endDesc}
	exit := space.Exit{Name: "end", Desc: "an exit that lends to the end", Location: &end}

	homeDesc := "Here is home"
	home := space.Location{Desc: homeDesc}
	homeExit := space.Exit{Name: "home", Desc: "a glowing tunnel leads home", Location: &home}

	start.AddExit(exit)
	start.AddExit(homeExit)

	player := Player{Name: "Tester", Location: &start}

	player.MoveToNamedExit("end")
	if player.Location.Desc != endDesc {
		t.Errorf("Expected players location to be %s but was %s", endDesc, player.Location.Desc)
	}

}

//Test to force safe error handling of attemp to move a non-existing location
func TestNonExistingMoveToFNamdeExit(t *testing.T) {

	startDesc := "This is where you start"

	start := space.Location{Desc: startDesc}

	endDesc := "This is where you end"
	end := space.Location{Desc: endDesc}
	exit := space.Exit{Name: "end", Desc: "an exit that lends to the end", Location: &end}

	homeDesc := "Here is home"
	home := space.Location{Desc: homeDesc}
	homeExit := space.Exit{Name: "home", Desc: "a glowing tunnel leads home", Location: &home}

	start.AddExit(exit)
	start.AddExit(homeExit)

	player := Player{Name: "Tester", Location: &start}

	err := player.MoveToNamedExit("nowhere")
	if err == nil {
		t.Errorf("Expected err when moving to nowhere but nil")
	}
}

func TestMovePlayer(t *testing.T) {

	startDesc := "This is where you start"
	start := space.Location{Desc: startDesc}
	player := Player{Name: "Tester", Location: &start}

	playerLocations := new(PlayerMap)
	playerLocations.PL = make([]PlayerLocation, 0, 100)
	pl := PlayerLocation{&player, &start}
	playerLocations.PL = append(playerLocations.PL, pl)

	playerLocations.MovePlayer(player.Location, &player)
	players := playerLocations.GetPlayersAtLocation(player.Location)

	p0name := (*players)[0].Name
	if p0name != "Tester" {

		t.Errorf("Expected %s but was %s", player.Name, p0name)

	}

	homeDesc := "Here is home"
	home := space.Location{Desc: homeDesc}
	player.Location = &home

	playerLocations.MovePlayer(player.Location, &player)
	players = playerLocations.GetPlayersAtLocation(&start)
	if players != nil {

		t.Errorf("Expected nil but was %s who thinks they should be at %s but is at %s", (*players)[0].Name, player.Location.Desc, start.Desc)

	}

}

func TestLookWithOtherPlayers(t *testing.T) {
	playerLocations := new(PlayerMap)
	playerLocations.PL = make([]PlayerLocation, 0, 100)

	startDesc := "This is where you start"
	start := space.Location{Desc: startDesc}
	player1 := Player{Name: "Tester1", Location: &start}
	playerLocations.MovePlayer(player1.Location, &player1)
	player2 := Player{Name: "Tester2", Location: &start}
	playerLocations.MovePlayer(player2.Location, &player2)
	player3 := Player{Name: "Tester3", Location: &start}
	playerLocations.MovePlayer(player3.Location, &player3)

	result := player1.Look(playerLocations)

	expected := startDesc + "\n" + "Tester2 is here.\nTester3 is here.\n"
	if result != expected {
		t.Errorf("Expected\n %s\nbut was\n%s\n", expected, result)
	}

}

func TestGoRoutineMoves(t *testing.T) {
	playerLocations := new(PlayerMap)

	startDesc := "This is where you start"
	start := space.Location{Desc: startDesc}
	player1 := Player{Name: "Tester1", Location: &start}
	player2 := Player{Name: "Tester2", Location: &start}
	player3 := Player{Name: "Tester3", Location: &start}

	go playerLocations.MovePlayer(player3.Location, &player3)
	go playerLocations.MovePlayer(player2.Location, &player2)
	go playerLocations.MovePlayer(player1.Location, &player1)
	time.Sleep(5000)

	result := player1.Look(playerLocations)

	expected := startDesc + "\n" + "Tester2 is here.\nTester3 is here.\n"
	if result != expected {
		t.Errorf("Expected\n %s\nbut was\n%s\n", expected, result)
	}

}
