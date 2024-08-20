package apis

import (
	"fmt"
	"net/http"
)

func CreateFile(path, fileSystemName, addr string, port int) {
	url := "http://" + addr + ":" + fmt.Sprint(port) + "/api/v1/meta/createFile"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("path", path)
	req.Header.Add("fileSystemName", fileSystemName)
	client := &http.Client{}
	client.Do(req)
}
