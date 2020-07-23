package wechat_bot

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Markdown struct {
	Content string `json:"content"`
}

type MsgContent struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

func SendMsg(url string, content MsgContent) error {
	contentJson, _ := json.Marshal(content)
	body := strings.NewReader(string(contentJson))
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
