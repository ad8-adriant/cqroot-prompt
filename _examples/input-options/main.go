package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ad8-adriant/cqroot-prompt"
	"github.com/ad8-adriant/cqroot-prompt/input"
)

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func main() {
	val, err := prompt.New().Ask("Input with Help:").Input(
		"Blah blah",
		input.WithHelp(true),
		input.WithWidth(5),
		input.WithCharLimit(5),
	)
	CheckErr(err)

	fmt.Printf("{ %s }\n", val)
}
