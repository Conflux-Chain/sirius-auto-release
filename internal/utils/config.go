package utils

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"fmt"
)

func GetPortBindingForConfig(globalConfig *config.Global, proxyConfig *config.Proxy, containerConfig *config.Container) []string {
	ports := []string{}

	if globalConfig.Space == config.ALL_SPACE {
		ports = append(ports, fmt.Sprintf("%d:%d", containerConfig.CoreSpace.Port, proxyConfig.CoreSpace.Port))
		ports = append(ports, fmt.Sprintf("%d:%d", containerConfig.ESpace.Port, proxyConfig.ESpace.Port))
	} else if globalConfig.Space == config.CORE_SPACE {
		ports = append(ports, fmt.Sprintf("%d:%d", containerConfig.CoreSpace.Port, proxyConfig.CoreSpace.Port))
	} else if globalConfig.Space == config.E_SPACE {
		ports = append(ports, fmt.Sprintf("%d:%d", containerConfig.ESpace.Port, proxyConfig.ESpace.Port))
	} else {
		return nil
	}

	return ports
}
