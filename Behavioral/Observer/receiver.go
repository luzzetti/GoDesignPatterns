package main

import "fmt"

type publisher interface {
	registerObserver(obs subscriber)
	unregisterObserver(obs subscriber)
	notifyAll()
}

type DataReceiver struct {
	Name string
}

func (l *DataReceiver) onUpdate(data string) {
	fmt.Println("Listener:", l.Name, "got data:", data)
}
