package main

// Adapter to bridge the client 'television' interface, with the legacy tv

type brandLegacyAdapter struct {
	legacyTv *BrandLegacyTv
}

func (t brandLegacyAdapter) turnOn() {
	t.legacyTv.setOnState(true)
}

func (t brandLegacyAdapter) turnOff() {
	t.legacyTv.setOnState(false)

}

func (t brandLegacyAdapter) volumeUp() int {
	vol := t.legacyTv.getVolume() + 1
	t.legacyTv.setVolume(vol)
	return vol
}

func (t brandLegacyAdapter) volumeDown() int {
	vol := t.legacyTv.getVolume() - 1
	t.legacyTv.setVolume(vol)
	return vol
}

func (t brandLegacyAdapter) channelUp() int {
	ch := t.legacyTv.getCurrentChan() + 1
	t.legacyTv.setChannel(ch)
	return ch
}

func (t brandLegacyAdapter) channelDown() int {
	ch := t.legacyTv.getCurrentChan() - 1
	t.legacyTv.setChannel(ch)
	return ch
}

func (t brandLegacyAdapter) goToChannel(channel int) {
	t.legacyTv.setChannel(channel)
}
