package client_test

import (
	"github.com/geekdada/flomo-cli/client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMemo_Submit_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"Test"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	defer server.Close()

	memo := client.Memo{
		Content: "Test",
		Api:     server.URL,
	}

	message, err := memo.Submit(false)

	assert.Equal(t, "记录成功", *message)
	assert.Nil(t, err)
}

func TestMemo_Submit_WithTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"content":"Test\n\n#TestTag"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":0,"message":"记录成功"}`))
	}))

	defer server.Close()

	memo := client.Memo{
		Content: "Test",
		Tag:     "TestTag",
		Api:     server.URL,
	}

	message, err := memo.Submit(false)

	assert.Equal(t, "记录成功", *message)
	assert.Nil(t, err)
}

func TestMemo_Submit_LackOfArgs(t *testing.T) {
	func() {
		memo := client.Memo{
			Api:     "",
			Content: "Test",
		}

		_, err := memo.Submit(false)

		assert.EqualError(t, err, "lack of necessary arguments")
	}()

	func() {
		memo := client.Memo{
			Api:     "Test",
			Content: "",
		}

		_, err := memo.Submit(false)

		assert.EqualError(t, err, "lack of necessary arguments")
	}()
}

func TestMemo_Submit_BadServer(t *testing.T) {
	memo := client.Memo{
		Api:     "http://server_not_exist",
		Content: "Test",
	}

	_, err := memo.Submit(false)

	assert.Equal(t, strings.HasPrefix(err.Error(), "failed to make a request to server"), true)
}

func TestMemo_Submit_BadResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{`))
	}))

	defer server.Close()

	memo := client.Memo{
		Content: "Test",
		Api:     server.URL,
	}

	_, err := memo.Submit(false)

	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestMemo_Submit_500(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
		rw.Write([]byte(`{}`))
	}))

	defer server.Close()

	memo := client.Memo{
		Content: "Test",
		Api:     server.URL,
	}

	_, err := memo.Submit(false)

	assert.Equal(t, "status 500: err response error", err.Error())
}

func TestMemo_Submit_400(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(400)
		rw.Write([]byte(`{}`))
	}))

	defer server.Close()

	memo := client.Memo{
		Content: "Test",
		Api:     server.URL,
	}

	_, err := memo.Submit(false)

	assert.Equal(t, "status 400: err request is not valid", err.Error())
}