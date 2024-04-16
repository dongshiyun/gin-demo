package consul

import (
	"fmt"
	"gd/config"

	"github.com/hashicorp/consul/api"
)

// https://developer.hashicorp.com/consul/docs
// https://developer.hashicorp.com/consul/api-docs/agent/service#register-service
func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = config.Address

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           config.CheckHttp,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

// 根据ID获取服务
func GetServiceByID(id string) {
	cfg := api.DefaultConfig()
	cfg.Address = config.Address

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s", id))
	if err != nil {
		panic(err)
	}

	svrHost := ""
	svrPort := 0
	for key, value := range data {
		svrHost = value.Address
		svrPort = value.Port
		fmt.Println("获取到服务：" + key)
		fmt.Println("IP:PORT", svrHost, svrPort)

	}
}
