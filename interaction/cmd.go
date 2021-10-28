package interaction

import (
	"gomud/entity"
	"strings"
)

//tokenise using whitespace a string that is also converted to lower case
func Tokenise(command string) []string {

	command = strings.ToLower(command)

	tokens := strings.Fields(command)

	return tokens

}

func Do(cmdTokens []string, player *entity.Player) string {

	result := "unknown command"
	if cmdTokens[0] == "look" {

		result = player.Look()

	}
	if cmdTokens[0] == "go" {
		err := player.MoveToNamedExit(cmdTokens[1])
		if err != nil {
			result = err.Error()
		} else {
			result = player.Look()
		}
	}
	return result

}
