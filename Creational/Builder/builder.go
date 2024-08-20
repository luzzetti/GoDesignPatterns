package main

import (
	"fmt"
)

type NotificationBuilder struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.title = title
}

func (nb *NotificationBuilder) SetSubtitle(subtitle string) {
	nb.subtitle = subtitle
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.priority = priority
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	/*
		Return any error from here. To avoid sending a lot of _, err while
		using the builder
	*/

	if nb.icon != "" && nb.subtitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle when using an icon")
	}
	if nb.priority > 5 {
		return nil, fmt.Errorf("you need to specify a priority between 0 and 5")
	}

	return &Notification{
		title:    nb.title,
		subtitle: nb.subtitle,
		message:  nb.message,
		image:    nb.image,
		icon:     nb.icon,
		priority: nb.priority,
		notType:  nb.notType,
	}, nil

}
