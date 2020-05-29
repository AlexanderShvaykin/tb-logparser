package tb_logparser

import (
	"fmt"
)

type printer interface {
	Write(line string)
}

type stdoutPrinter struct{}

func (s stdoutPrinter) Write(line string) {
	fmt.Println(line)
}
