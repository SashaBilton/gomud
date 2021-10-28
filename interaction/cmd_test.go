package interaction

import (
	"fmt"
	"gomud/entity"
	"gomud/space"
	"testing"
)

//Test the seperation of commands into tokens
func TestTokenise(t *testing.T) {

	command := "Go west"

	token := Tokenise(command)
	if token[0] != "go" || token[1] != "west" {
		t.Errorf("Expected go and west tokens but was %s amd %s", token[0], token[1])
	}
}

//The look command should return a string containing the description and exists a location has
func TestLookCommand(t *testing.T) {

	command := "LOOK"
	commandTokens := Tokenise(command)
	player := setupSimpleWorld()
	result := do(commandTokens, &player)

	expected := fmt.Sprintf("This is where you start\nan exit that lends to the end\na glowing tunnel leads home\n")

	if result != expected {
		t.Errorf("Expected \n%s but was %s", expected, result)
	}

}

//The go command should move the player to the location pointed to by the named exit
func TestGoCommand(t *testing.T) {

	command := "gO end"
	commandTokens := Tokenise(command)
	player := setupSimpleWorld()
	result := do(commandTokens, &player)

	expected := fmt.Sprintf("This is where you end\n")

	if result != expected {
		t.Errorf("Expected \n%s but was %s", expected, result)
	}

}

func setupSimpleWorld() entity.Player {
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

	player := entity.Player{Name: "Tester", Location: &start}
	return player
}
