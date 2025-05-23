package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Name        string
	Description string
}

func getTasks(file *os.File) []Task {
	var tasks []Task

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatal("failed to parse json!")
	}
	return tasks
}

func runOperation[T any](operate func(*os.File) T) T {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal("failed to open json file!")
	}
	defer file.Close()

	return operate(file)
}

func main() {
	tasks := runOperation(getTasks)
	for _, task := range tasks {
		fmt.Println(task)
	}
}
