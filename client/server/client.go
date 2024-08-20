package server

import (
	"bytes"
	"fmt"
	"github.com/huiming23344/mindfs/client/config"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

type Registry struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type clientServer struct {
	ServiceName string
	Addr        string
	Port        int
	Registry    Registry
}

var ClientServer clientServer

func InitClient() {
	cfg := config.GlobalConfig()
	addrs, err := getHostIPAddresses()
	if err != nil {
		fmt.Printf("get host ip failed: %v\n", err)
		return
	}
	ClientServer = clientServer{
		ServiceName: cfg.Client.ServiceName,
		Addr:        addrs[len(addrs)-1],
		Port:        cfg.Client.Port,
		Registry: Registry{
			Address: cfg.Registry.Address,
			Port:    cfg.Registry.Port,
		},
	}
}

func splitFile() {
	// 定义文件路径和分块大小
	filePath := "./file.txt"
	chunkSize := 1024 * 1024 // 1KB

	// 打开原始文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建子文件夹，如果它不存在
	subDir := "chunks"
	err = os.MkdirAll(subDir, 0755)
	if err != nil {
		panic(err)
	}

	// 遍历文件并分块存储
	chunkCounter := 0
	buffer := new(bytes.Buffer)
	for {
		// 读取 1KB 的数据块
		chunk := make([]byte, chunkSize)
		n, err := file.Read(chunk)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break // 到达文件末尾
		}

		// 写入数据块到 buffer
		buffer.Write(chunk[:n])

		// 创建子文件路径
		chunkPath := subDir + "/chunk-" + fmt.Sprintf("%05d", chunkCounter) + ".dat"

		// 创建子文件并写入数据块
		chunkFile, err := os.Create(chunkPath)
		if err != nil {
			panic(err)
		}
		defer chunkFile.Close()
		_, err = chunkFile.Write(chunk[:n])
		if err != nil {
			panic(err)
		}

		// 增加块计数器
		chunkCounter++
	}

	// 关闭原始文件
	file.Close()

	// 输出块计数器以确认所有块都已处理
	fmt.Println("All chunks have been processed and stored.")
}

func reassembleFromChunks(chunkDir string, outputFilePath string) error {
	// 创建输出文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 遍历子文件夹中的所有文件
	files, err := os.ReadDir(chunkDir)
	if err != nil {
		return err
	}

	// 按顺序读取文件块
	for _, file := range files {
		// 创建子文件路径
		chunkPath := filepath.Join(chunkDir, file.Name())

		// 打开子文件
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}
		defer func(chunkFile *os.File) {
			err := chunkFile.Close()
			if err != nil {
				fmt.Println("Error closing file:", err)
			}
		}(chunkFile)

		// 读取子文件内容
		chunkBytes, err := ioutil.ReadAll(chunkFile)
		if err != nil {
			return err
		}

		// 写入子文件内容到输出文件
		_, err = outputFile.Write(chunkBytes)
		if err != nil {
			return err
		}
	}

	// 返回成功
	return nil
}

func getHostIPAddresses() ([]string, error) {
	var addresses []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 遍历网络接口
	for _, i := range interfaces {
		// 获取接口的地址列表
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		// 遍历地址列表
		for _, addr := range addrs {
			// 检查是否为IPv4地址
			ip := addr.(*net.IPNet)
			if ip.IP.To4() != nil {
				// 添加IPv4地址到结果列表
				addresses = append(addresses, ip.IP.String())
			}
		}
	}

	return addresses, nil
}
