package entity

import (
	"gomud/space"
)

type Player struct {
	Name     string
	Location *space.Location
}

func (player *Player) MoveToNamedExit(name string) {
	player.Location = player.Location.Exits.Location
}
