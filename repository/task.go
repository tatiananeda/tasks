package repository

import (
	"github.com/google/uuid"
	"log"
	"time"
)

const Layout = "2006-Jan-02"

type Task struct {
	Title, Description string
	Due                time.Time
	IsComplete         bool
	Id                 string
}

func NewTask(title, description, due, id string) *Task {
	if title == "" || due == "" {
		log.Fatal("Title and Date due are required")
	}

	uuid := id

	if id == "" {
		uuid = genUUID()
	}

	t := Task{
		Title:       title,
		Description: description,
		Due:         getDueDate(due),
		Id:          uuid,
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
		t.Due = getDueDate(due)
	}
}

func (t *Task) ToggleComplete() {
	t.IsComplete = !t.IsComplete
}

func genUUID() string {
	uuid, error := uuid.NewRandom()
	ExitIfError(error, "Failed to generate UUID")
	return uuid.String()
}

func getDueDate(due string) time.Time {
	d, err := time.Parse(Layout, due)
	ExitIfError(err, "Wrong date format "+Layout+" expected")
	return d
}
