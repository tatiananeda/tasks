package main

import (
	repo "github.com/tatiananeda/tasks/repository"
)

func main() {
	r := repo.Repository{
		List: make([]*repo.Task, 0),
	}

	r.AddTask(repo.NewTask("Title", "Do Homework", "2024-Oct-29"))
	r.AddTask(repo.NewTask("Title2", "Do Smth", "2024-Oct-28"))
	r.AddTask(repo.NewTask("No title", "", "2024-Oct-30"))
	r.AddTask(repo.NewTask("Repo", "Commit homework to git", "2024-Oct-28"))

	r.EditTask("Title2", "Title3", "", "2024-Oct-29")
	r.DeleteTask("Title3")

	r.ChangeCompleteStatus("Title")
	r.ChangeCompleteStatus("No title")

	r.PrintTasks(r.GetFiltered(true))
}
