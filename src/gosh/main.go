package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strconv"
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
		dateCmd := exec.Command("date")

		dateOut, err := dateCmd.StdoutPipe()
		if err != nil {
			panic(err)
		}

		dateCmd.Start()

		dateBytes, err := ioutil.ReadAll(dateOut)
		if err != nil {
			panic(err)
		}

		// This will reap the child
		//dateCmd.Wait()

		fmt.Printf("% 4d >> date: %s", i, dateBytes)
	}

	//time.Sleep(60 * time.Second)
}
