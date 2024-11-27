package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ad8-adriant/cqroot-prompt"
	"github.com/ad8-adriant/cqroot-prompt/write"
)

var ErrContentTooShort = errors.New("content too short")

func validateLength(content string) error {
	if len(content) < 10 {
		return ErrContentTooShort
	} else {
		return nil
	}
}

func main() {
	val, err := prompt.New().Ask("Write something:").
		Write("Blah Blah Blah Blah", write.WithValidateFunc(validateLength))
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else if errors.Is(err, ErrContentTooShort) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
