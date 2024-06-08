package main

import (
	"scheduler/process"
	"scheduler/scheduler"
)

func main() {

	//example
	processes := []*process.Process {
		{Pid: 1, Priority: 5, CpuTime: 10, IoOperations: 3, RemainingTime: 10},
		
		{Pid: 2, Priority: 3, CpuTime: 5, IoOperations: 1, RemainingTime: 5},
		{Pid: 3, Priority: 1, CpuTime: 8, IoOperations: 2, RemainingTime: 8},
	}
	
	scheduler.FeedbackScheduler(processes)
}
