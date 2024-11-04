package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	// "io/fs"
	"os"
)

const filename = "./data.json"

type Repository struct {
	List []*Task
}

func (r *Repository) AddTask(t *Task) {
	r.List = append(r.List, t)
}

func (r *Repository) Save() {
	res, err := json.Marshal(r.List)
	ExitIfError(err, "Error marshaling to JSON")
	errWrite := os.WriteFile(filename, res, 0644)
	fmt.Println(errWrite)
	ExitIfError(errWrite, "Error writing file")
}

func (r *Repository) ViewTask(id string) {
	for _, task := range r.List {
		if task.Id == id {
			fmt.Printf("Task:\n %s \n %s \n Due: %s", task.Title, task.Description, task.Due)
			return
		}
	}
	fmt.Printf("\nNo task with id %s found", id)
}

func (r *Repository) EditTask(id string, t, d, due string) {
	for _, task := range r.List {
		if task.Id == id {
			task.Edit(t, d, due)
			return
		}
	}
	fmt.Printf("\nNo task with id %s found", id)
}

func (r *Repository) DeleteTask(id string) {
	for i, task := range r.List {
		if task.Id == id {
			task = nil
			r.List = append(r.List[:i], r.List[i+1:]...)
			return
		}
	}
	fmt.Printf("\nNo task with id %s found", id)
}

func (r *Repository) PrintTasks(tasks []*Task) {
	fmt.Println("\nTasks:")
	for _, task := range tasks {
		fmt.Printf("Title:%s\nDescription:%s\nDue: %s\nIsCompleted: %t\nId: %s\n", task.Title, task.Description, task.Due.Format("01-02-2006"), task.IsComplete, task.Id)
	}
}

func (r *Repository) ChangeCompleteStatus(id string) {
	for _, task := range r.List {
		if task.Id == id {
			task.ToggleComplete()
			return
		}
	}
	fmt.Printf("\nNo task with id %s found", id)
}

func (r *Repository) GetFiltered(complete bool) []*Task {
	res := make([]*Task, 0)
	for _, t := range r.List {
		if t.IsComplete == complete {
			res = append(res, t)
		}
	}
	return res
}

func Rehidrate() Repository {
	r := Repository{
		List: make([]*Task, 0),
	}

	f, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return r
		}
	}

	var i []Task
	errUnmashal := json.Unmarshal(f, &i)
	ExitIfError(errUnmashal, "Error decoding data")

	for _, val := range i {
		r.AddTask(&val)
	}

	return r
}
