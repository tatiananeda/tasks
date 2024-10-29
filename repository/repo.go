package repository

import (
	"fmt"
)

type Repository struct {
	List []*Task
}

func (r *Repository) AddTask(t *Task) {
	r.List = append(r.List, t)
}

func (r *Repository) ViewTask(t string) {
	for _, task := range r.List {
		if task.Title == t {
			fmt.Printf("Task:\n %s \n %s \n Due: %s", task.Title, task.Description, task.Due.Format("01-02-2006"))
			return
		}
	}
	fmt.Printf("\nNo task with title %s found", t)
}

func (r *Repository) EditTask(title string, t, d, due string) {
	for _, task := range r.List {
		if task.Title == title {
			task.Edit(t, d, due)
			return
		}
	}
	fmt.Printf("\nNo task with title %s found", title)
}

func (r *Repository) DeleteTask(title string) {
	for i, task := range r.List {
		if task.Title == title {
			task = nil
			r.List = append(r.List[:i], r.List[i+1:]...)
			return
		}
	}
	fmt.Printf("\nNo task with title %s found", title)
}

func (r *Repository) PrintTasks(tasks []*Task) {
	fmt.Println("\nTasks:")
	for _, task := range tasks {
		fmt.Printf("Title:%s\nDescription:%s\nDue: %s\nIsCompleted: %t\n\n", task.Title, task.Description, task.Due.Format("01-02-2006"), task.IsComplete)
	}
}

func (r *Repository) ChangeCompleteStatus(title string) {
	for _, task := range r.List {
		if task.Title == title {
			task.ToggleComplete()
			return
		}
	}
	fmt.Printf("\nNo task with title %s found", title)
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
