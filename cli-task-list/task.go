package main

type Priority int

const (
	PriorityHigh = iota
	PriorityNormal
	PriorityLow
)

var priorityName = map[Priority]string {
	PriorityHigh = "High"
	PriorityNormal = "Normal"
	PriorityLow = "Low"
}

type Task struct {
	text string
	checked bool
	priority 
}
