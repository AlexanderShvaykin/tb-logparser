package tb_logparser

import (
	"bufio"
	"fmt"
	sync_buff "github.com/my/repo/pkg/sync-buff"
	"github.com/my/repo/pkg/tree"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

func remoteGrep(pattern string, server string, printer printer) {
	defer wg.Done()
	var out sync_buff.Buff
	defer out.Reset()

	cmd := exec.Command("ssh", server, "-t", "grep", pattern, LogName)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		errors = append(errors, fmt.Sprint(err))
	}
	res := out.String()
	if res != "" {
		printer.Write(res)
	}
}

func localGrep(pattern string, printer printer) {
	var wg sync.WaitGroup
	var osFiles []*os.File
	defer func() {
		for _, file := range osFiles {
			_ = file.Close()
		}
	}()

	paths, err := tree.ReadTree(".")
	if err != nil {
		panic(err)
	}
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		osFiles = append(osFiles, file)
		lines := make(chan string, 1)
		wg.Add(1)
		go func(out chan<- string, f *os.File) {
			defer wg.Done()

			scanner := bufio.NewScanner(f)
			buf := make([]byte, 0, 64*1024)
			scanner.Buffer(buf, 1024*1024)
			for scanner.Scan() {
				text := strings.TrimSpace(scanner.Text())
				re := regexp.MustCompile(pattern)
				if re.MatchString(text) {
					out <- text
				}
			}
			if scanner.Err() != nil {
				log.Fatalf("error: %s\n", scanner.Err())
			}
			close(out)
		}(lines, file)

		wg.Add(1)
		go func(in chan string) {
			defer wg.Done()

			for line := range in {
				printer.Write(line)
			}
		}(lines)
	}
	wg.Wait()
}
