package config

import (
	"match/lib/file"
)

var ServiceConfig *serviceConfig

type serviceConfig struct {
	RoomServiceMaxCount int `json:"room_service_max_count"`
	RoomPlayerMaxCount  int `json:"room_player_max_count"`
}

var serviceConfigs map[string]*serviceConfig

func initServiceConfig(datePath string) {
	serviceConfigs = make(map[string]*serviceConfig)
	err := file.LoadJsonToObject(datePath, &serviceConfigs)
	if err != nil {
		panic(err)
	}
	for k, config := range serviceConfigs {
		if k == ServiceType {
			ServiceConfig = config
		}
	}
}
