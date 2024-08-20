package apis

import (
	"encoding/json"
	"fmt"
	"github.com/huiming23344/mindfs/client/server"
	"io"
	"log"
	"net/http"
)

type Service struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}

func Discovery(service string) []Service {
	url := fmt.Sprintf("http://%s:%d/api/discovery", server.ClientServer.Registry.Address, server.ClientServer.Registry.Port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return nil
	}
	if service != "" {
		req.Header.Set("service", service)
	}
	client := &http.Client{}
	log.Println("Heartbeat to registry...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return nil
	}
	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return nil
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(rspBody))
	var services []Service
	err = json.Unmarshal(rspBody, &services)
	if err != nil {
		log.Println("Error unmarshalling response body: ", err)
		return nil
	}
	return services
}
