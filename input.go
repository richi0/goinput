package input

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

func isPipeInput() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}

func isFlagInput() bool {
	return len(strings.Join(flag.Args(), "")) != 0
}

func isValidInput(isPipe bool, isFlag bool) bool {
	if !isPipe && !isFlag {
		return false
	}
	if isPipe && isFlag {
		panic("pipe and flag input")
	}
	return true
}

func GetInput() string {
	var res string
	flag.Parse()
	isPipe := isPipeInput()
	isFlag := isFlagInput()
	if !isValidInput(isPipe, isFlag) {
		return res
	}
	if !isPipe && isFlag {
		res = strings.Join(flag.Args(), " ")
	} else {
		reader := bufio.NewReader(os.Stdin)
		res, _ = reader.ReadString(0)
	}
	return res
}
