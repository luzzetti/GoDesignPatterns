package main

import (
	"fmt"
	"sync"
)

type Mylogger struct {
	loglevel int
}

func (l *Mylogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

func (l *Mylogger) SetLogLevel(level int) {
	l.loglevel = level
}

// The logger instance
var logger *Mylogger

// Thread-safety
var once sync.Once

// This is NOT thread-safe
func getLogger() *Mylogger {
	if logger == nil {
		fmt.Println("Creating new logger")
		logger = &Mylogger{}
	}
	fmt.Println("Returning Logger Singleton")
	return logger
}

func getThreadSafeLogger() *Mylogger {
	once.Do(func() {
		fmt.Println("Creating new thread-safe logger")
		logger = &Mylogger{}
	})

	fmt.Println("Returning thread-safe Logger Singleton")
	return logger
}
