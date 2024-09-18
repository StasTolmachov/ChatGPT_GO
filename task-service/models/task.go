package models

type Task struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Done        bool   `json:"Done"`
}
