package v1

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDiscovery(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		discoveryAPI()
	}
}

func discoveryAPI() {
	url := "http://127.0.0.1:8180/api/discovery"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("serviceName", "time-service")
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
