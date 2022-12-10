package publish

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func PublishSignUp(user uint) {
	signUpObject := registerObj{
		ToUser: user,
		Type:   "register",
	}

	client := &http.Client{}

	json, err := json.Marshal(signUpObject)
	if err != nil {
		log.Println(err)
	}
	req, errReq := http.NewRequest(http.MethodPut, os.Getenv("CLOUD_RELAY_PUB"), bytes.NewBuffer(json))
	if errReq != nil {
		log.Println(errReq)
	}
	req.Header.Set("topic", "email")
	req.Header.Set("service", "Profile Service")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, errSend := client.Do(req)
	if errSend != nil {
		log.Println(errSend)
	}
}
