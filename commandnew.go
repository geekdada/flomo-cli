package main

import (
	"github.com/geekdada/flomo-cli/client"
	"log"
	"os"
)

type NewCommand struct {
	Verbose bool `short:"V" long:"verbose" description:"Show verbose debug information"`

	Api string `long:"api" description:"flomo API address" required:"true"`

	Content string `short:"c" long:"content" description:"Content to be sent" required:"true"`

	Tag string `short:"t" long:"tag" description:"Additional tag"`
}

func (x *NewCommand) Execute(args []string) error {
	memo := client.Memo{
		Content: x.Content,
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
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println(*responseMessage)
	os.Exit(0)

	return nil
}
