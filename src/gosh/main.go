package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strconv"
	"github.com/kardianos/osext"
	"path/filepath"
	"path"
	"bufio"
	"bytes"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must have an iteration count")
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to convert arg: %s", err)
	}

	for i := 0; i < count; i++ {
		phantomPath, err := exec.LookPath("phantomjs")
		if err != nil {
			log.Fatal("Could not find phantomjs in PATH")
		}

		filename, _ := osext.Executable()
		workDir := filepath.Dir(filename)

		testJs := path.Join(workDir, "test.js")

		cmd := exec.Command(phantomPath, testJs)

		cmdOut, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}

		cmd.Start()

		dateBytes, err := ioutil.ReadAll(cmdOut)
		if err != nil {
			panic(err)
		}

		// This will reap the child
		//dateCmd.Wait()

		scanner := bufio.NewScanner(bytes.NewReader(dateBytes))
		for scanner.Scan() {
			fmt.Printf("% 4d >> output: %s\n", i, scanner.Text())
		}
	}

	syscall.Kill(os.Getpid(), syscall.SIGKILL)
}
