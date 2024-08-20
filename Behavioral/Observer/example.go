package main

func main() {

	listener1 := DataReceiver{Name: "AA"}
	listener2 := DataReceiver{Name: "BB"}
	listener3 := DataReceiver{Name: "CC"}

	subj := &DataEmitter{}

	subj.registerReceiver(listener1)
	subj.registerReceiver(listener2)
	subj.registerReceiver(listener3)

	subj.ChangeItem("Monday!")
	subj.ChangeItem("Friday!")
	subj.ChangeItem("Sunday!")

	subj.unregisterReceiver(listener1)
	subj.unregisterReceiver(listener2)
	subj.unregisterReceiver(listener3)

	subj.ChangeItem("Another message")

}
