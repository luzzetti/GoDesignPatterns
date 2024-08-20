package main

import "fmt"

type BrandLegacyTv struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *BrandLegacyTv) setOnState(on bool) {
	if on {
		fmt.Println("Legacy TV is on")
	} else {
		fmt.Println("Legacy TV is off")
	}
	tv.tvOn = on
}

func (tv *BrandLegacyTv) getVolume() int {
	return tv.currentVolume
}

func (tv *BrandLegacyTv) setVolume(vol int) {
	fmt.Println("Legacy TV volume set to", vol)
	tv.currentVolume = vol
}

func (tv *BrandLegacyTv) setChannel(channel int) {
	fmt.Println("Legacy TV channel set to", channel)
	tv.currentChan = channel
}

func (tv *BrandLegacyTv) getCurrentChan() int {
	return tv.currentChan
}

func (tv *BrandLegacyTv) getTvOn() bool {
	return tv.tvOn
}
