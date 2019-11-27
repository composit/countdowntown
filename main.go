package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	// TODO make these configurable
	doneMsg        = "did"
	interval       = 5 * time.Second // seconds
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

	f, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := updateFile(mins, f); err != nil {
		log.Fatal(err)
	}

	if err := countdown(mins, f); err != nil {
		log.Fatal(err)
	}
}

func updateFile(mins int, f *os.File) error {
	b := []byte(fmt.Sprintf("%d", mins))

	l, err := f.WriteAt(b, 0)
	if err != nil {
		return err
	}

	if err := f.Truncate(int64(l)); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}

func countdown(mins int, f *os.File) error {
	d := time.Duration(mins) * time.Minute

	ticker := time.NewTicker(interval)

timer:
	for {
		select {
		case <-ticker.C:
			d = d - interval
			fmt.Printf("left: %s\n", d)
			fmt.Printf("left divided: %#v\n", d/time.Minute)
			fmt.Printf("left inted: %d\n", int(d/time.Minute))
			left := int(d / time.Minute)

			if _, err := f.WriteAt([]byte(fmt.Sprintf("%d", left)), 0); err != nil {
				return err
			}

			if int(d) < 1 {
				break timer
			}
		}
	}

	if _, err := f.WriteAt([]byte(doneMsg), 0); err != nil {
		return err
	}

	return nil
}
