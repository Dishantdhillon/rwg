package main

import (
	"flag"
	"fmt"

	"github.com/roz3x/rwg/html"
	"github.com/roz3x/rwg/passage"
)

var (
	count     = flag.Int("v", 10, "give the max length of word")
	passagein = flag.Int("p", 10, "give the no of words")
	typeout   = flag.Int("t", 0, "takes the type [0-2]")
	useHTML   = flag.Int("html", 0, "if set 1 then gives html output")
)

func main() {
	flag.Parse()

	if *useHTML == 1 {
		text, err := html.Generate(*passagein)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			fmt.Printf("%v", text)
		}
	} else {
		text, err := passage.Generate(*passagein, *count, *typeout)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			fmt.Printf("%v", text)
		}

	}
	fmt.Printf("\n")
}
