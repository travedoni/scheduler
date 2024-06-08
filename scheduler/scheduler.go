package scheduler

import (
	"container/heap"
	"fmt"
	"scheduler/process"
)

func adjustPriority(p *process.Process) {
	if p.IoOperations > 2 {
		p.Priority -= 1
	}

	if p.CpuTime > 7 {
		p.Priority += 1
	}

	if p.Priority < 1 {
		p.Priority = 1
	}

}

func FeedbackScheduler(processes []*process.Process) {
	pq := NewPriorityQueue()
	for _, p := range processes {
		heap.Push(pq, p)
	}
	
	time := 0
	timeSlice := 1
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*process.Process)
		executionTime := timeSlice
		if p.RemainingTime < timeSlice {
			executionTime = p.RemainingTime
		}

		p.RemainingTime -= executionTime
		time += executionTime

		p.TotalTime = time
		p.WaitTime = time - p.CpuTime
		fmt.Printf("Process %d executed w/ priority %d\n",
			p.Pid, p.Priority)

		adjustPriority(p)

		// Put back the process into the priority queue
		// if there is time left
		if p.RemainingTime > 0 {
			heap.Push(pq, p)
		}
	}

	// Print process stats
	for _, p := range processes {
		fmt.Printf("Process %d, wait time: %d, total time: %d\n",
			p.Pid, p.WaitTime, p.TotalTime)
	}
}

