package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/huiming23344/mindfs/metaServer/server"
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

func Register() {
	body := Service{
		ServiceName: server.MetaServer.ServiceName,
		ServiceId:   server.MetaServer.ServiceId,
		IpAddress:   server.MetaServer.Addr,
		Port:        server.MetaServer.Port,
	}
	jsonData, _ := json.Marshal(body)
	reqBody := bytes.NewBuffer(jsonData)
	url := fmt.Sprintf("http://%s:%d/api/register", server.MetaServer.Registry.Address, server.MetaServer.Registry.Port)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
	log.Println("Heartbeat to registry...")
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

func Unregister() {
	body := Service{
		ServiceName: server.MetaServer.ServiceName,
		ServiceId:   server.MetaServer.ServiceId,
		IpAddress:   server.MetaServer.Addr,
		Port:        server.MetaServer.Port,
	}
	jsonData, _ := json.Marshal(body)
	reqBody := bytes.NewBuffer(jsonData)
	url := fmt.Sprintf("http://%s:%d/api/unregister", server.MetaServer.Registry.Address, server.MetaServer.Registry.Port)
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return
	}
	client := &http.Client{}
	log.Println("Heartbeat to registry...")
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
