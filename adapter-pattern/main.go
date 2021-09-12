package main

import (
	"design-pattern/adapter-pattern/client"
	"design-pattern/adapter-pattern/provider"
	"design-pattern/adapter-pattern/service"
)

func main() {

	c := &client.Client{}
	mac := &service.Mac{}

	c.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &provider.Windows{}
	windowsMachineAdapter := &service.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	c.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
