package main

import "fmt"

type BrandValidTv struct {
	channel int
	vol     int
	isOn    bool
}

func (t BrandValidTv) turnOn() {
	fmt.Println("Valid TV is on")
	t.isOn = true
}

func (t BrandValidTv) turnOff() {
	fmt.Println("Valid TV is off")
	t.isOn = false
}

func (t BrandValidTv) volumeUp() int {
	t.vol++
	fmt.Println("Valid TV volume set to", t.vol)
	return t.vol
}

func (t BrandValidTv) volumeDown() int {
	fmt.Println("Valid TV volume set to", t.vol)
	t.vol--
	return t.vol
}

func (t BrandValidTv) channelUp() int {
	fmt.Println("Valid TV channel set to", t.channel)
	t.channel++
	return t.channel
}

func (t BrandValidTv) channelDown() int {
	fmt.Println("Valid TV channel set to", t.channel)
	t.channel--
	return t.channel
}

func (t BrandValidTv) goToChannel(channel int) {
	fmt.Println("Valid TV channel set to", channel)
	t.channel = channel
}
