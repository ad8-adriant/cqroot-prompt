package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ad8-adriant/cqroot-prompt"
	"github.com/ad8-adriant/cqroot-prompt/write"
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
	val, err := prompt.New().Ask("Write:").Write(
		"Blah blah blah...",
		write.WithHelp(true),
		write.WithCharLimit(800),
		write.WithWidth(20),
		write.WithLineNumbers(true),
	)
	CheckErr(err)

	fmt.Println(val)
}
