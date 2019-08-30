package utils

import "fmt"

// FakeProcess contains the information about trully procecess only to simulate the procecess functionality
type FakeProcess struct {
	Pid      int32
	Name     string
	Status   string // Status returns the process status. Return value could be one of these. R: Running S: Sleep T: Stop I: Idle Z: Zombie W: Wait L: Lock The character is same within all supported platforms.
	User     string
	CPU      float64
	Memory   float32
	Priority int32
}

// ToArray converts the process information to an array to process the information
func (proc *FakeProcess) ToArray() []string {
	pid := fmt.Sprintf("%d", proc.Pid)
	cpu := fmt.Sprintf("%f", proc.CPU)
	memory := fmt.Sprintf("%f", proc.Memory)
	priority := fmt.Sprintf("%d", proc.Priority)
	return []string{
		pid,
		proc.Name,
		proc.Status,
		proc.User,
		cpu,
		memory,
		priority,
	}
}
