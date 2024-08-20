package main

func main() {

	// write a for range from 0 to 10
	for i := 0; i < 10; i++ {
		go getThreadSafeLogger()
	}

	log := getThreadSafeLogger()
	log.SetLogLevel(1)
	log.Log("This is a log message")

	log = getThreadSafeLogger()
	log.SetLogLevel(2)
	log.Log("This is another log message")

	log = getThreadSafeLogger()
	log.SetLogLevel(3)
	log.Log("This is another another log message")
}
