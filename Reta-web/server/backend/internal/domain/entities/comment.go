package entities

import "time"

type Comment struct {
	UUID     string    `json:"UUID"`
	PostID   string    `json:"post_id"`
	Author   string    `json:"author"`
	Message  string    `json:"message"`
	MediaURL []string  `json:"medial_urls"`
	Time     time.Time `json:"time"`
}
