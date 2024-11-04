package main

import (
	repo "github.com/tatiananeda/tasks/repository"
)

func main() {
	r := repo.Rehidrate()

	r.AddTask(repo.NewTask("Title", "Do Homework", "2024-Oct-29", ""))
	r.AddTask(repo.NewTask("Title2", "Do Smth", "2024-Oct-28", ""))
	r.AddTask(repo.NewTask("No title", "", "2024-Oct-30", ""))
	r.ChangeCompleteStatus("ac87521e-2bfd-4853-a2bd-b0460ff63acb")
	r.EditTask("ac87521e-2bfd-4853-a2bd-b0460ff63acb", "Changed", "", "")

	r.PrintTasks(r.List)
	r.Save()
}
