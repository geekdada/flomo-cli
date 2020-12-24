package main

import (
	"fmt"
	"github.com/geekdada/flomo-cli/client"
	"github.com/pkg/errors"
	"log"
	"os"
)

type NewCommand struct {
	Verbose bool `short:"V" long:"verbose" description:"Show verbose debug information"`

	Api string `long:"api" description:"flomo API address"`

	Content string `short:"c" long:"content" description:"Content to be sent" required:"true"`

	Tag string `short:"t" long:"tag" description:"Additional tag"`
}

func (x *NewCommand) Execute(args []string) error {
	api := os.Getenv("FLOMO_API")

	if api == "" {
		api = x.Api
	}

	if api == "" {
		return errors.New("you must specify flomo API address either using FLOMO_API env or --api")
	}

	memo := client.Memo{
		Content: x.Content,
		Tag:     x.Tag,
		Api:     api,
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
