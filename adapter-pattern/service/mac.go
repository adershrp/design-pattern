package service

import "fmt"

// Mac no adapter required
type Mac struct {
}

// InsertIntoLightningPort implementing the product
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
