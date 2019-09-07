package parallel

import (
	"fmt"
	"strings"
	"time"

	"github.com/erickmx/process_visualizer/utils"
)

// ParallelizeProcess is the function to run in the go routine and act as a process execution simulation
func ParallelizeProcess(process *utils.FakeProcess, running int32) {
	// fmt.Println("ParallelizeProcess", process.Pid, running)
	if strings.ContainsAny(process.Status, "Ss") && running == process.Pid {
		process.Status = "R"
		fmt.Println("Running process ", process)
		time.Sleep(3 * time.Second)
		process.Status = "S"
	}
}
