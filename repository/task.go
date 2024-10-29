package repository

import (
	"log"
	"time"
)

const Layout = "2006-Jan-02"

type Task struct {
	Title, Description string
	Due                time.Time
	IsComplete         bool
}

func NewTask(title, description, due string) *Task {
	if title == "" || due == "" {
		log.Fatal("Title and Date due are required")
	}
	d, err := time.Parse(Layout, due)
	if err != nil {
		log.Fatal("Wrong date format " + Layout + " expected")
	}
	t := Task{
		Title:       title,
		Description: description,
		Due:         d,
	}
	return &t
}

func (t *Task) Edit(title, description, due string) {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if due != "" {
		d, err := time.Parse(Layout, due)
		if err != nil {
			log.Fatal("Wrong date format " + Layout + " expected")
		}
		t.Due = d
	}
}

func (t *Task) ToggleComplete() {
	t.IsComplete = !t.IsComplete
}
