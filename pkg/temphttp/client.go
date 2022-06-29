package temphttp

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string, client *http.Client) *Client {
	return &Client{baseURL: baseURL, client: client}
}

func (c *Client) Convert(choice string, temp int) (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/temp-converter", nil)
	q := req.URL.Query()

	q.Add("choice", fmt.Sprint(choice))
	q.Add("temp", fmt.Sprintf("%v", temp))
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	s := string(body)
	return s, nil
}
