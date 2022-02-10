package compatibility

import (
	"github.com/erikgeiser/promptkit/confirmation"
)

// Prompts user with yes/no about whether changes are backwards compatible
//
// Bool return value represents backwards compatibility: if `false`, changes are "BREAKING" and not backwards compatible.
func Prompt() (bool, error) {
	prompt := confirmation.New("Is this change backwards compatible?", confirmation.Yes)
	return prompt.RunPrompt()
}
