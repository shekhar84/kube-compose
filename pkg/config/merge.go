package config

import (
	"fmt"
)

// extendedService is merged into service
func merge(service *Service, extendedService *Service) {
	// rules based on https://docs.docker.com/compose/extends/#adding-and-overriding-configuration
	mergeStringMap(service.Environment, extendedService.Environment)
	mergePortBindings(service.Ports, extendedService.Ports)
	// TODO https://github.com/jbrekelmans/kube-compose/issues/48
}

func mergeStringMap(intoStringMap map[string]string, fromStringMap map[string]string) {
	for k, v := range fromStringMap {
		if _, ok := intoStringMap[k]; !ok {
			intoStringMap[k] = v
		}
	}
}

func mergePortBindings(intoPorts []PortBinding, fromPorts []PortBinding) {
	fmt.Println(intoPorts)
	fmt.Println(fromPorts)
	intoPorts = append(intoPorts, fromPorts...)
	fmt.Println(intoPorts)
}
