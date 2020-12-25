package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCommand_Usage(t *testing.T) {
	newCommand := NewCommand{}

	newCommand.Usage()
}

func TestNewCommand_Run(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"测试内容"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	newCommand := NewCommand{
		Api: server.URL,
	}

	if _, err := newCommand.Run([]string{"测试内容"}); err != nil {
		log.Fatal(err)
	}
}
