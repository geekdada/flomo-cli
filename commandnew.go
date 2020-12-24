package main

import (
	"bufio"
	"fmt"
	"github.com/geekdada/flomo-cli/client"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
)

type NewCommand struct {
	Verbose bool `short:"V" long:"verbose" description:"Show verbose debug information"`

	Api string `long:"api" description:"flomo API address" env:"FLOMO_API"`

	Tag string `short:"t" long:"tag" description:"Additional tag"`
}

func (x *NewCommand) Usage() string {
	return "[new command options] <memo content>"
}

func (x *NewCommand) Execute(args []string) error {
	var content string

	if isInputFromPipe() {
		content = getStdinContent(os.Stdin)
	} else {
		content = strings.Join(args, " ")
	}

	if content == "" {
		return errors.New("you must specify the content of the memo")
	}

	memo := client.Memo{
		Content: content,
		Tag:     x.Tag,
		Api:     x.Api,
	}

	responseMessage, err := memo.Submit(x.Verbose)

	if err != nil {
		switch err.(type) {
		case *client.ResponseError:
			re, _ := err.(*client.ResponseError)

			if re.StatusCode >= 400 && re.StatusCode < 500 {
				os.Exit(2)
			} else {
				os.Exit(1)
			}
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}

	log.Println(*responseMessage)
	os.Exit(0)

	return nil
}

func isInputFromPipe() bool {
    fileInfo, _ := os.Stdin.Stat()
    return fileInfo.Mode() & os.ModeCharDevice == 0
}

func getStdinContent(r io.Reader) string {
	var runes []rune
	var output string

	reader := bufio.NewReader(r)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		runes = append(runes, input)
	}

	for j := 0; j < len(runes); j++ {
		output += fmt.Sprintf("%c", runes[j])
	}

	return output
}
