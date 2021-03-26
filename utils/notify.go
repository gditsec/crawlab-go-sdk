package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gditsec/crawlab-go-sdk/datapool"
	"github.com/gditsec/crawlab-go-sdk/entity"
)

type Response struct {
	Status  string `json:"status`
	Message string `json:"message"`
}

func NotifyTarget(item entity.Item, files []string) error {
	target := datapool.GetTargetConfig()

	record, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	data := entity.Item{
		"files":  files,
		"record": string(record),
	}

	response := &Response{}
	if err := Post(target.Notify, data, response); err != nil {
		return err
	}
	if response.Status != "ok" {
		return errors.New("notify target error: " + response.Message)
	}
	return nil
}

func Post(url string, data interface{}, ret interface{}) error {
	client := &http.Client{Timeout: 5 * time.Second}
	body, _ := json.Marshal(data)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(ret)
}
