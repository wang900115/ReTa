package entities

import "time"

type Comment struct {
	UUID    string    `json:"UUID"`
	Author  string    `json:"author"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
