package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ad8-adriant/cqroot-prompt"
	"github.com/ad8-adriant/cqroot-prompt/input"
)

func main() {
	val, err := prompt.New().Ask("Input an number:").
		Input("", input.WithInputMode(input.InputNumber))
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
