package comm

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTPClient Http Client wrapper struct
type HTTPClient struct {
	BaseURL string
}

// CreateHTTPClient : Creates new HTTPClient instance
func CreateHTTPClient(url string) HTTPClient {
	c := HTTPClient{url}
	return c
}

// Get Makes a GET request to the given path
func (c HTTPClient) Get(path string, platform string, lang string) (response []byte, err error) {
	req, err := http.NewRequest("GET", c.BaseURL+platform+path, nil)
	if err != nil {
		log.Fatal("Error reading request", err)
	}

	req.Header.Set("Accept-Language", lang)

	client := &http.Client{Timeout: time.Second * 10}

	log.Print("Sending http request", req)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error reading reponse. %s", err)
	}
	defer resp.Body.Close()

	log.Print("http response", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading response", err)
	}
	return body, err
}
