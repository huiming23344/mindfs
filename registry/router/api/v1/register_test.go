package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {

	body1 := ServiceReq{
		ServiceName: "time-service",
		ServiceId:   "time-service-1",
		IpAddress:   "127.0.0.1",
		Port:        8280,
	}
	body2 := ServiceReq{
		ServiceName: "time-service",
		ServiceId:   "time-service-2",
		IpAddress:   "127.0.0.1",
		Port:        8281,
	}
	register(body1)
	register(body2)
}

func register(body ServiceReq) {
	url := "http://127.0.0.1:8180/api/register"
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return
	}
	repBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(repBody))
}
