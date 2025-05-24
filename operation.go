package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const filename string = "tasks.json"

func runOperation(operate func(*os.File)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open json file")
	}
	defer file.Close()

	operate(file)
}

func displayTasks(file *os.File) {

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatal("failed to parse json data")
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
}

func (newTask *Task) addTasks(file *os.File) {

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatal("failed to parse json file")
	}

	newTask.ID = uuid.New()
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()
	tasks = append(tasks, *newTask)

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal("failed to update tasks")
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Fatal("failed to update tasks")
	}
}
