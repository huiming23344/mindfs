package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DataServer struct {
	Id   string
	Ip   string
	Port int
}

type WriteDataReq struct {
	data        string
	dataServers []*DataServer
}

func WriteData(data string, dataServer DataServer) error {
	reqbody := WriteDataReq{
		data:        data,
		dataServers: []*DataServer{&dataServer},
	}
	body, err := json.Marshal(reqbody)
	if err != nil {
		return err
	}
	bodyreader := bytes.NewReader(body)
	// Send data to the data server
	req, err := http.NewRequest("POST", "http://"+dataServer.Ip+":"+fmt.Sprint(dataServer.Port)+"/api/v1/writeData", bodyreader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
