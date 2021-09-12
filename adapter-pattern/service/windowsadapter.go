package service

import (
	"fmt"

	"design-pattern/adapter-pattern/provider"
)

// WindowsAdapter constructor injection
type WindowsAdapter struct {
	WindowMachine *provider.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
