package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/huiming23344/mindfs/metaServer/config"
	"github.com/huiming23344/mindfs/metaServer/meta"
	"net"
)

type Registry struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type metaServer struct {
	ServiceName string
	ServiceId   string
	Addr        string
	Port        int
	Registry    Registry
	Users       map[string]*meta.User
	Groups      map[string]*meta.UserGroup
	Dir         *meta.Directory
}

var MetaServer metaServer

func InitServer() {
	cfg := config.GlobalConfig()
	addrs, err := getHostIPAddresses()
	if err != nil {
		fmt.Printf("get host ip failed: %v\n", err)
		return
	}
	serviceId := uuid.New().String()
	MetaServer = metaServer{
		ServiceName: cfg.Server.ServiceName,
		ServiceId:   serviceId,
		Addr:        addrs[len(addrs)-1],
		Port:        cfg.Server.Port,
		Registry: Registry{
			Address: cfg.Registry.Address,
			Port:    cfg.Registry.Port,
		},
		Users:  make(map[string]*meta.User),
		Groups: make(map[string]*meta.UserGroup),
		Dir:    meta.NewDirectory("/"),
	}
	AddUser("admin", "admin")
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

func AddUser(username, password string) error {
	if _, exists := MetaServer.Users[username]; exists {
		return fmt.Errorf("user '%s' already exists", username)
	}
	MetaServer.Users[username] = &meta.User{
		Username: username,
		Password: password,
	}
	return nil
}

func DeleteUser(username string) error {
	if _, exists := MetaServer.Users[username]; !exists {
		return fmt.Errorf("user '%s' not exists", username)
	}
	delete(MetaServer.Users, username)
	return nil
}

func ListUsers() []string {
	users := make([]string, 0, len(MetaServer.Users))
	for username := range MetaServer.Users {
		users = append(users, username)
	}
	return users
}

func AddGroup(groupName string) error {
	if _, exists := MetaServer.Groups[groupName]; exists {
		return fmt.Errorf("group '%s' already exists", groupName)
	}
	MetaServer.Groups[groupName] = &meta.UserGroup{
		Name: groupName,
	}
	return nil
}

func DeleteGroup(groupName string) error {
	if _, exists := MetaServer.Groups[groupName]; !exists {
		return fmt.Errorf("group '%s' not exists", groupName)
	}
	delete(MetaServer.Groups, groupName)
	return nil
}

func AddUserToGroup(username, groupName string) error {
	user, exists := MetaServer.Users[username]
	if !exists {
		return fmt.Errorf("user '%s' not exists", username)
	}
	group, exists := MetaServer.Groups[groupName]
	if !exists {
		return fmt.Errorf("group '%s' not exists", groupName)
	}
	group.Users = append(group.Users, user)
	return nil
}

func RemoveUserFromGroup(username, groupName string) error {
	group, exists := MetaServer.Groups[groupName]
	if !exists {
		return fmt.Errorf("group '%s' not exists", groupName)
	}
	for i, user := range group.Users {
		if user.Username == username {
			group.Users = append(group.Users[:i], group.Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user '%s' not in group '%s'", username, groupName)
}
