package main

import "fmt"

type subscriber interface {
	onUpdate(data string)
}

type DataEmitter struct {
	receivers []DataReceiver
	field     string
}

func (d *DataEmitter) ChangeItem(data string) {
	d.field = data
	d.notifyAll()
}

func (d *DataEmitter) registerReceiver(aReceiver DataReceiver) {
	fmt.Println("Receiver registered:", aReceiver.Name)
	d.receivers = append(d.receivers, aReceiver)
}

func (d *DataEmitter) unregisterReceiver(o DataReceiver) {
	fmt.Println("Listener unregistered:", o.Name)
	var newobs []DataReceiver
	for _, obs := range d.receivers {
		if obs.Name != o.Name {
			newobs = append(newobs, obs)
		}
	}
	d.receivers = newobs

}

func (d *DataEmitter) notifyAll() {
	for _, obs := range d.receivers {
		obs.onUpdate(d.field)
	}
}
