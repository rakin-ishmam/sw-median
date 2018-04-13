package cmd

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/rakin-ishmam/sw-median/fread"
	"github.com/rakin-ishmam/sw-median/median"
)

func calcMedian(fileLoc string, wsize int) {
	// t := time.Now()

	f, err := fread.NewCsv(fileLoc)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	m := median.New(wsize)
	re := regexp.MustCompile("[0-9]+")

	for {
		f.Next()
		if f.Err() == io.EOF {
			break
		}
		if f.Err() != nil {
			log.Fatal(f.Err())
		}

		v := f.Data()[0]
		d, err := strconv.Atoi(re.FindString(v))

		if err != nil {
			log.Fatal("convertion", err)
		}

		m.AddDelay(d)

		fmt.Printf("%v\r\n", m.GetMedian())
	}

	// d := time.Since(t)
	// fmt.Println("tot time =", d.Nanoseconds())
}
