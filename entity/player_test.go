package entity

import (
	"gomud/space"
	"testing"
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

	start.AddExit(&exit)
	start.AddExit(&homeExit)

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

	start.AddExit(&exit)
	start.AddExit(&homeExit)

	player := Player{Name: "Tester", Location: &start}

	err := player.MoveToNamedExit("nowhere")
	if err == nil {
		t.Errorf("Expected err when moving to nowhere but nil")
	}
}
