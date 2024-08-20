package main

import "fmt"

/***
This is the interface used by the 'client' application.
It has to be adapted to two different television models
with different methods and APIs
*/

type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnOn()
	turnOff()
	goToChannel(ch int)
}

func main() {

	// A TV already compliant with the client interface
	var tv2 television
	tv2 = &BrandValidTv{
		channel: 9,
		vol:     20,
		isOn:    true,
	}

	fmt.Println("--------------------")

	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(68)
	tv2.turnOff()
	fmt.Println("--------------------")

	// a TV with older/legacy interface, that needs to be adapted
	tv1 := &BrandLegacyTv{
		currentChan:   35,
		currentVolume: 14,
		tvOn:          true,
	}

	var adapter television
	adapter = &brandLegacyAdapter{legacyTv: tv1}

	adapter.turnOn()
	adapter.volumeUp()
	adapter.volumeDown()
	adapter.channelUp()
	adapter.channelDown()
	adapter.goToChannel(68)
	adapter.turnOff()

	fmt.Println("--------------------")

}
