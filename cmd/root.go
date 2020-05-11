package cmd

import (
	"fmt"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

func Execute() {
	var out writer

	wg.Add(1)
	go grep("360882015971112437", "/webapps/teachbase/shared/log/logstasher.log", &out)

	wg.Add(1)
	go grep("360882015971112437", "/webapps/teachbase/shared/log/logstasher.log", &out)
	wg.Wait()
}

func grep(pattern string, path string, out *writer) {
	defer wg.Done()

	cmd := exec.Command("ssh", "yc-g1", "-t", "grep", pattern, path)
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type writer struct{}

func (w *writer) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
