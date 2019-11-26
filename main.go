package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide a number of minutes. For example: cdt 10")
	}

	mins, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := countdown(mins); err != nil {
		log.Fatal(err)
	}
}

func countdown(mins int) error {
	d := time.Duration(mins) * time.Minute
	i := 5 * time.Second

	ticker := time.NewTicker(i)

	f, err := os.Create("/home/matt/.cdt")

	if _, err = f.WriteAt([]byte(fmt.Sprintf("%d", mins)), 0); err != nil {
		return err
	}

timer:
	for {
		select {
		case <-ticker.C:
			d = d - i
			left := int(d / time.Minute)

			if _, err = f.WriteAt([]byte(fmt.Sprintf("%d", left)), 0); err != nil {
				return err
			}

			if err := f.Sync(); err != nil {
				return err
			}

			if int(d) < 1 {
				break timer
			}
		}
	}

	if _, err = f.WriteAt([]byte("did"), 0); err != nil {
		return err
	}

	return nil
}
