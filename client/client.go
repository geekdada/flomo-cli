package client

import (
	"encoding/json"
	"fmt"
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Memo struct {
	Content string
	Tag string
	Api string
}

type Payload struct {
	Content string `json:"content"`
}

func (m *Memo) Submit(verbose bool)  (*string, error) {
	content := strings.TrimSpace(m.Content)

	if m.Api == "" || content == "" {
		return nil, errors.New("lack of necessary arguments")
	}

	if m.Tag != "" {
		content += fmt.Sprintf("\n\n#%s", m.Tag)
	}

	timeout := 3000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	payloadJSON, _ := json.Marshal(Payload{
		content,
	})
	body := ioutil.NopCloser(strings.NewReader(string(payloadJSON)))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	if verbose {
		log.Printf("Raw content: %s", content)
		log.Printf("Payload JSON: %s", payloadJSON)
	}

	response, err := client.Post(m.Api, body, headers)

	if err != nil {
		return nil, errors.Wrap(err, "failed to make a request to server")
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var responseData map[string]interface{}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return nil, err
	}

	if verbose {
		log.Printf("Response Body: %v", responseData)
	}

	statusCode := response.StatusCode

	if statusCode >= 200 && statusCode < 400 {
		message := responseData["message"].(string)

		return &message, nil
	} else if statusCode >= 400 && statusCode < 500 {
		return nil, &ResponseError{
			Err:        errors.New("request is not valid"),
			StatusCode: statusCode,
		}
	} else {
		return nil, &ResponseError{
			Err:        errors.New("response error"),
			StatusCode: statusCode,
		}
	}
}
