package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	// TODO make these configurable
	interval       = 1 * time.Minute
	outputFilePath = "/home/matt/.cdt"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a number of minutes. For example: countdowntown 10")
	}

	mins, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := updateFile(mins, outputFilePath); err != nil {
		log.Fatal(err)
	}

	if err := countdown(mins, outputFilePath); err != nil {
		log.Fatal(err)
	}

	if err := writeDoneMsg(mins, outputFilePath); err != nil {
		log.Fatal(err)
	}
}

func updateFile(mins int, filePath string) error {
	b := []byte(fmt.Sprintf("%d", mins))

	if err := ioutil.WriteFile(filePath, b, 0644); err != nil {
		return err
	}

	return nil
}

func countdown(mins int, filePath string) error {
	d := time.Duration(mins) * time.Minute

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			d = d - interval
			remaining := int(d / time.Minute)

			if err := updateFile(remaining, filePath); err != nil {
				return err
			}

			if d < 1 {
				return nil
			}
		}
	}
}

func writeDoneMsg(mins int, filePath string) error {
	msg := fmt.Sprintf("did %d", mins)
	if err := ioutil.WriteFile(filePath, []byte(msg), 0644); err != nil {
		return err
	}

	return nil
}
