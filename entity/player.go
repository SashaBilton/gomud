package entity

import (
	"errors"
	"gomud/space"
)

type Player struct {
	Name     string
	Location *space.Location
}

func (player *Player) MoveToNamedExit(name string) error {

	exit := player.Location.Exits[name]
	if exit == nil {
		return errors.New("not found")
	} else {
		player.Location = exit.Location
		return nil
	}
}

func (player *Player) Look() string {
	result := player.Location.Desc + "\n"

	for _, exit := range player.Location.Exits {
		result += exit.Desc + "\n"
	}
	return result
}
