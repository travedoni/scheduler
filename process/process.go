package process

// Process is a single process
type Process struct {
	Pid		int
	Priority	int
	CpuTime		int
	IoOperations	int
	WaitTime	int
	TotalTime 	int
	RemainingTime	int
}

// NewProcess creates a new process
func NewProcess(pid, priority, cpuTime, ioOperations int) *Process {
	return &Process {
		Pid:		pid,
		Priority:	priority,
		CpuTime: 	cpuTime,
		IoOperations: 	ioOperations,
		WaitTime: 	0,
		TotalTime:	0,
		RemainingTime: cpuTime,
	}
}

// Execute simulates the execution of the process for a time slice
func (p *Process) Execute(timeSlice int) {
	if p.RemainingTime > timeSlice {
		p.RemainingTime -= timeSlice
	} else {
		p.RemainingTime = 0
	}
}

// UpdateWaitingTime updates the waiting time for the process
func (p *Process) UpdateWaitingTime(waitTime int) {
	p.WaitTime += waitTime
}

// UpdateTotalTime updates the total time for the process
func (p *Process) UpdateTotalTime(totalTime int) {
	p.TotalTime = totalTime
}
