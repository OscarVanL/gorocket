package gorocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HookMessage struct {
	Text        string           `json:"text"`
	Attachments []HookAttachment `json:"attachments"`
}

type HookAttachment struct {
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
	ImageURL  string `json:"image_url"`
	Color     string `json:"color"`
}

type HookResponse struct {
	Success bool `json:"success"`
}

func (c *Client) Hooks(msg *HookMessage, token string) (*HookResponse, error) {
	opt, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/hooks/%s", c.baseURL, token),
		bytes.NewBuffer(opt))

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		log.Fatal("Request error")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp := HookResponse{}

	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
