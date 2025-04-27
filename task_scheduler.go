package main

import (
	"fmt"
)

type Task struct {
	ID     int
	Name   string
	IsDone bool
}

type Scheduler struct {
	Tasks  []Task
	nextID int
}

func (self *Scheduler) CompleteTask(id int) bool {
	for _, task := range self.Tasks {
		if task.ID == id {
			self.Tasks[id].IsDone = true
			return true
		}
	}
	return false
}

func (self *Scheduler) ListPendingTasks() []Task {
	var sliceOfTasks []Task
	for _, task := range self.Tasks {
		if !task.IsDone {
			sliceOfTasks = append(sliceOfTasks, task)
		}
	}
	return sliceOfTasks
}

func (self *Scheduler) ListCompletedTasks() []Task {
	var sliceOfTasks []Task
	for _, task := range self.Tasks {
		if task.IsDone {
			sliceOfTasks = append(sliceOfTasks, task)
		}
	}
	return sliceOfTasks
}

func (self *Scheduler) AddTask(name string) {
	if name == "" {
		fmt.Println("task is empty adding \"default\" to task name")
		name = "default"
	}
	task := Task{ID: self.nextID, Name: name, IsDone: false}
	self.nextID++
	self.Tasks = append(self.Tasks, task)
	fmt.Println("Task Added : ", task.Name)
}

func main() {
	var Scheduler Scheduler
	Scheduler.AddTask("Go to work")
	Scheduler.AddTask("workout")
	Scheduler.AddTask("Check Scores")
	Scheduler.AddTask("Make tutorial video about Docker")
	Scheduler.AddTask("Publish tutorial video")

	fmt.Println(Scheduler.CompleteTask(1))
	fmt.Println(Scheduler.CompleteTask(3))
	fmt.Println(Scheduler.ListPendingTasks())
	fmt.Println(Scheduler.ListCompletedTasks())
}
