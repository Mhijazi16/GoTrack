package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
