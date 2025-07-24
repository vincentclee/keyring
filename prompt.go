package keyring

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// PromptFunc is a function used to prompt the user for a password.
type PromptFunc func(string) (string, error)

// TerminalPrompt prompts the user for input on the terminal, hiding the input for password security.
func TerminalPrompt(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	b, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(b), nil
}

// FixedStringPrompt returns a PromptFunc that always returns the provided fixed value.
func FixedStringPrompt(value string) PromptFunc {
	return func(_ string) (string, error) {
		return value, nil
	}
}
