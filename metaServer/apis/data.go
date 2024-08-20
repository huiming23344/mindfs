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

type CopyDataReq struct {
	dataServer DataServer
}

func CopyData(inodeId string, dataServer DataServer) error {
	reqbody := CopyDataReq{
		dataServer: dataServer,
	}
	body, err := json.Marshal(reqbody)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)
	// Send data to the data server
	req, err := http.NewRequest("POST", "http://"+dataServer.Ip+":"+fmt.Sprint(dataServer.Port)+"/api/v1/copyData", bodyReader)
	if err != nil {
		return err
	}
	req.Header.Add("inodeId", inodeId)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
