package publish

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type registerObj struct {
	ToUser  uint
	Type    string
	Message []byte
}

type codeObj struct {
	Code string
}

func PublishRegister(user uint, signUpCode string) {

	code := codeObj{
		Code: signUpCode,
	}
	jsonCode, errCode := json.Marshal(code)
	if errCode != nil {
		log.Println(errCode)
	}

	registerObj := registerObj{
		ToUser:  user,
		Type:    "registerCode",
		Message: jsonCode,
	}

	client := &http.Client{}

	json, err := json.Marshal(registerObj)
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
