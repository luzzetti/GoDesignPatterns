package main

import "fmt"

func main() {

	var bldr = newNotificationBuilder()

	bldr.SetTitle("New Notification")
	bldr.SetIcon("anIcon")
	bldr.SetSubtitle("a Subtitle")
	bldr.SetPriority(3)
	// missing fields

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
		return
	}

	fmt.Printf("Created notification: %+v ", notif)

}
