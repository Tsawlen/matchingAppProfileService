package publish

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type logMessage struct {
	Message string
}

func PublishLogger(message string) {
	logMessage := logMessage{
		Message: message,
	}

	client := &http.Client{}

	json, err := json.Marshal(logMessage)
	if err != nil {
		log.Println(err)
	}
	req, errReq := http.NewRequest(http.MethodPut, os.Getenv("CLOUD_RELAY_PUB"), bytes.NewBuffer(json))
	if errReq != nil {
		log.Println(errReq)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, errSend := client.Do(req)
	if errSend != nil {
		log.Println(errSend)
	}
}
