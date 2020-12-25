package main

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

var update = flag.Bool("update", false, "update test files with results")

func TestCLI_New(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"测试内容"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	defer server.Close()

	if err := exec.Command("go", "run", ".", "new", "--api", server.URL, "测试内容").Run(); err != nil {
		log.Fatal(err)
	}
}

func TestCLI_New_WithTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"测试内容\n\n#test"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	defer server.Close()

	if err := exec.Command("go", "run", ".", "new", "--api", server.URL, "--tag", "test", "测试内容").Run(); err != nil {
		log.Fatal(err)
	}
}

func TestCLI_New_Pipe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"测试内容"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	defer server.Close()

	command1 := exec.Command("echo", "测试内容")
	command2 := exec.Command("go", "run", ".", "new", "--api", server.URL)

	r, w := io.Pipe()
	command1.Stdout = w
	command2.Stdin = r

	var command2Buffer bytes.Buffer
	command2.Stdout = &command2Buffer

	logFatal(command1.Start())
	logFatal(command2.Start())
	logFatal(command1.Wait())
	logFatal(w.Close())
	logFatal(command2.Wait())
	logFatal(func() error {
		_, err := io.Copy(os.Stdout, &command2Buffer)
		return err
	}())
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
