package apis

import (
	"fmt"
	"github.com/huiming23344/mindfs/registry/server"
	"net/http"
)

func Invalid() {
	metaSrv := server.GetService("metaServer")
	url := "http://" + metaSrv.IpAddress + ":" + fmt.Sprint(metaSrv.Port) + "/api/invalid"
	req, err := http.NewRequest("POST", url, nil)
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
	defer resp.Body.Close()
}
