package main

type Task struct {
	Name        string
	Description string
	Status      string
}

func main() {
	runOperation(displayTasks)
}
