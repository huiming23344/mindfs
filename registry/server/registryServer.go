package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/huiming23344/mindfs/registry/config"
	"net"
	"time"
)

type ServiceReq struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}

type Service struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
	Timer       int    `json:"timer"`
}

type registryServer struct {
	IpAddress      string
	Port           int
	ServiceMap     map[string]*Service
	ServiceNameMap map[string][]string
	ServiceCounter map[string]int
}

var RegistryServer registryServer

func InitRegistryServer() {
	cfg := config.GlobalConfig()
	serviceId := uuid.New().String()
	addrs, err := getHostIPAddresses()
	if err != nil {
		fmt.Printf("get host ip failed: %v\n", err)
		return
	}
	fmt.Printf("serviceId: %s\n", serviceId)
	RegistryServer = registryServer{
		IpAddress:      addrs[len(addrs)-1],
		Port:           cfg.Server.Port,
		ServiceMap:     make(map[string]*Service),
		ServiceNameMap: make(map[string][]string),
		ServiceCounter: make(map[string]int),
	}
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

func RegisterService(serviceReq ServiceReq) {
	fmt.Printf("register service: %v\n", serviceReq)
	cfg := config.GlobalConfig()
	service := &Service{
		Port:        serviceReq.Port,
		ServiceId:   serviceReq.ServiceId,
		ServiceName: serviceReq.ServiceName,
		IpAddress:   serviceReq.IpAddress,
		Timer:       cfg.Server.HeartBeatTimeout,
	}
	RegistryServer.ServiceMap[serviceReq.ServiceId] = service
	RegistryServer.ServiceNameMap[serviceReq.ServiceName] = append(RegistryServer.ServiceNameMap[serviceReq.ServiceName], serviceReq.ServiceId)

}

func CheckIsRight(serviceReq ServiceReq) bool {
	have := RegistryServer.ServiceMap[serviceReq.ServiceId]
	if have != nil {
		if have.ServiceName != serviceReq.ServiceName || have.IpAddress != serviceReq.IpAddress || have.Port != serviceReq.Port {
			return false
		}
		return true
	}
	return false
}

func UnregisterService(serviceReq ServiceReq) bool {
	fmt.Printf("unregister service: %v\n", serviceReq)
	if CheckIsRight(serviceReq) == false {
		return false
	}
	if _, ok := RegistryServer.ServiceMap[serviceReq.ServiceId]; ok {
		delete(RegistryServer.ServiceMap, serviceReq.ServiceId)
	}
	if _, ok := RegistryServer.ServiceNameMap[serviceReq.ServiceName]; ok {
		for i, serviceId := range RegistryServer.ServiceNameMap[serviceReq.ServiceName] {
			if serviceId == serviceReq.ServiceId {
				RegistryServer.ServiceNameMap[serviceReq.ServiceName] = append(RegistryServer.ServiceNameMap[serviceReq.ServiceName][:i], RegistryServer.ServiceNameMap[serviceReq.ServiceName][i+1:]...)
				return true
			}
		}
	}
	return false
}

func GetService(serviceName string) *Service {
	if serviceList, ok := RegistryServer.ServiceNameMap[serviceName]; ok {
		counter := RegistryServer.ServiceCounter[serviceName]
		RegistryServer.ServiceCounter[serviceName] = ((counter + 1) % len(serviceList))
		serviceID := serviceList[counter]
		return RegistryServer.ServiceMap[serviceID]
	}
	return nil
}

func GetAllService() []*Service {
	var services []*Service
	for _, service := range RegistryServer.ServiceMap {
		services = append(services, service)
	}
	return services
}

func HeartbeatService(serviceId string) {
	if _, ok := RegistryServer.ServiceMap[serviceId]; ok {
		RegistryServer.ServiceMap[serviceId].Timer = config.GlobalConfig().Server.HeartBeatTimeout
	}
}

func CheckService() {
	for _, service := range RegistryServer.ServiceMap {
		service.Timer--
		fmt.Printf("serviceId: %s, timer: %d\n", service.ServiceId, service.Timer)
		if service.Timer <= 0 {
			have := RegistryServer.ServiceMap[service.ServiceId]
			if have != nil {
				UnregisterService(ServiceReq{
					ServiceName: service.ServiceName,
					ServiceId:   service.ServiceId,
					IpAddress:   service.IpAddress,
					Port:        service.Port,
				})
			}
		}
	}
}

func Check() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		CheckService()
	}
}
