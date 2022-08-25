package main

import (
	"fmt"
)

type Client struct {
}

func (c *Client) InsertLightningConnectIntoComputer(com Computer) {
	fmt.Println("Client insert Lightning connector into computer")
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

//有点像AOP
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB")
	w.windowMachine.insertIntoUSBPort()
}
