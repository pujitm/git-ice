package subject

import "github.com/erikgeiser/promptkit/textinput"

func Prompt() (string, error) {
	prompt := textinput.New("Enter subject:")
	return prompt.RunPrompt()
}
