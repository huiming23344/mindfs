package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/huiming23344/mindfs/dataServer/config"
	"github.com/huiming23344/mindfs/dataServer/db"
	"net"
)

type Registry struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type dataServer struct {
	ServiceName string
	ServiceId   string
	Addr        string
	Port        int
	db          db.DB
	Registry    Registry
}

var DataServer dataServer

func InitServer() {
	cfg := config.GlobalConfig()
	addrs, err := getHostIPAddresses()
	if err != nil {
		fmt.Printf("get host ip failed: %v\n", err)
		return
	}
	DB, err := db.NewDB("./nodes/node0", cfg.Server.CacheCap)
	if err != nil {
		return
	}
	serviceId := uuid.New().String()
	DataServer = dataServer{
		ServiceName: cfg.Server.ServiceName,
		ServiceId:   serviceId,
		Addr:        addrs[len(addrs)-1],
		Port:        cfg.Server.Port,
		db:          DB,
		Registry: Registry{
			Address: cfg.Registry.Address,
			Port:    cfg.Registry.Port,
		},
	}
}

func Write(key, value string) error {
	err := DataServer.db.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func Read(key string) (string, error) {
	data, err := DataServer.db.Get(key)
	if err != nil {
		return "", err
	}
	return data, nil
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
