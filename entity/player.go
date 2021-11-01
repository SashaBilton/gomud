package entity

import (
	"errors"
	"fmt"
	"gomud/space"
	"sync"
)

type Player struct {
	Name     string
	Location *space.Location
}

type PlayerLocation struct {
	Player   *Player
	Location *space.Location
}

type PlayerMap struct {
	mu sync.Mutex
	PL []PlayerLocation
}

func (player *Player) MoveToNamedExit(name string) error {

	exit, found := player.Location.Exits[name]
	if !found {
		return errors.New("not found")
	} else {
		player.Location = exit.Location
		return nil
	}
}

func (player *Player) Look(playerLocations *PlayerMap) string {
	result := player.Location.Desc + "\n"

	for _, exit := range player.Location.Exits {
		result += exit.Desc + "\n"
	}
	others := playerLocations.GetPlayersAtLocation(player.Location)
	for _, pl := range *others {
		if pl.Name != player.Name {
			result += pl.Name + " is here.\n"
		}
	}
	return result
}

func (playerLocations *PlayerMap) MovePlayer(location *space.Location, player *Player) {
	found := false
	playerLocations.mu.Lock()
	defer playerLocations.mu.Unlock()
	for i, pl := range playerLocations.PL {

		if pl.Player == player {
			playerLocations.PL[i].Location = location
			found = true
			fmt.Printf("found player %s set to %s\n", pl.Player.Name, pl.Location.Desc)
		}

	}
	if !found {
		newpl := PlayerLocation{Player: player, Location: location}

		playerLocations.PL = append(playerLocations.PL, newpl)
	}

}

func (playerLocations *PlayerMap) GetPlayersAtLocation(location *space.Location) *[]Player {

	result := new([]Player)
	for _, pl := range playerLocations.PL {
		if pl.Location == location {
			*result = append(*result, *pl.Player)
		}

	}
	return result
}

func List(playerLocations *PlayerMap) string {
	result := ""
	for _, pl := range playerLocations.PL {
		result += pl.Player.Name + " " + pl.Location.Desc + "\n"

	}
	return result

}
