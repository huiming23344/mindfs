package apis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func SendChunk(path, addr string, port int) error {
	// get the chunk by path
	buffer := new(bytes.Buffer)
	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening file: ", err)
		return errors.New("error opening file")
	}
	defer file.Close()
	_, err = io.Copy(buffer, file)
	// send the chunk to the server
	url := fmt.Sprintf("http://%s:%d/api/getDateTime", addr, port)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return errors.New("http.NewRequest failed")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return errors.New("client.Do failed")
	}
	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return errors.New("error reading response body")
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(rspBody))
	return nil
}

func GetChunk(path, addr string, port int) error {
	url := fmt.Sprintf("http://%s:%d/api/getDateTime", addr, port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest failed, err:", err)
		return errors.New("http.NewRequest failed")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do failed, err:", err)
		return errors.New("client.Do failed")
	}
	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return errors.New("error reading response body")
	}
	log.Println("Response status code:", resp.Status)
	log.Println("Response body:", string(rspBody))
	return nil
}
