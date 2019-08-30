package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

// GetProcecess returns all the procecess information fromthe system to simulate the functionality
func GetProcecess() (fakeProcecess []*FakeProcess) {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Hubo un error al obtener los procesos", err)
	}
	for _, processEl := range processes {
		name, _ := processEl.Name()
		memory, _ := processEl.MemoryPercent()
		status, _ := processEl.Status()
		user, _ := processEl.Username()
		cpu, _ := processEl.CPUPercent()
		priority, _ := processEl.Nice()
		fakeProcess := &FakeProcess{
			Pid:      processEl.Pid,
			Name:     name,
			Status:   status,
			User:     user,
			CPU:      cpu,
			Memory:   memory,
			Priority: priority,
		}
		fakeProcecess = append(fakeProcecess, fakeProcess)
	}
	return
}

// SaveProcecess saves all the procecess in a file to use the information later
func SaveProcecess(procecess []*FakeProcess) (okay bool) {
	file, _ := CreateFile("./procecess.csv")
	defer file.Close()
	for _, process := range procecess {
		data := fmt.Sprintf("%d,%s,%s,%s,%f,%f,%d\n", process.Pid, process.Name, process.Status, process.User, process.CPU, process.Memory, process.Priority)
		_, err := file.WriteString(data)
		if err != nil {
			return false
		}
	}
	return true
}

// ReadProcecess reads the procecess information stored in a csv file and puts them in a slice of *FakeProcess
func ReadProcecess(path string) (procecess []*FakeProcess) {
	data, err := ReadFile("./procecess.csv")
	if err != nil {
		return
	}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line != "" {
			procStr := strings.Split(line, ",")
			pid, _ := strconv.ParseInt(procStr[0], 10, 32)
			name := procStr[1]
			status := procStr[2]
			user := procStr[3]
			cpu, _ := strconv.ParseFloat(procStr[4], 64)
			memory, _ := strconv.ParseFloat(procStr[5], 32)
			priority, _ := strconv.ParseInt(procStr[6], 10, 32)
			fakeProcess := &FakeProcess{
				Pid:      int32(pid),
				Name:     name,
				Status:   status,
				User:     user,
				CPU:      cpu,
				Memory:   float32(memory),
				Priority: int32(priority),
			}
			procecess = append(procecess, fakeProcess)
		}
	}
	return
}
