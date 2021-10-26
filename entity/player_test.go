package entity

import (
	"gomud/space"
	"testing"
)

func TestMoveToNamedExit(t *testing.T) {

	startDesc := "This is where you start"

	start := space.Location{Desc: startDesc}

	endDesc := "This is where you end"
	end := space.Location{Desc: endDesc}
	exit := space.Exit{Name: "end", Desc: "an exit that lends to the end", Location: &end}

	start.Exits = &exit

	player := Player{Name: "Tester", Location: &start}

	player.MoveToNamedExit("end")
	if player.Location.Desc != endDesc {
		t.Errorf("Expected players location to be %s but was %s", endDesc, player.Location.Desc)
	}

}
