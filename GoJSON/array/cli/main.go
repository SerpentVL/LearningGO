package main

import (
	"flag"
	"fmt"

	"gitlab.echelon.lan/polygon/akvs/akvs_jam_analyze/akvs_jam2cws"
)

func main() {
	var bugsPathname string

	flag.StringVar(&bugsPathname, "b", "./bugs.json", "bugs database")
	flag.Parse()

	if len(flag.Args()) < 2 {
		return
	}

	bugs, err := akvs_jam2cws.ReadBugs(bugsPathname)

	if err != nil {
		panic(err)
	}

	fmt.Printf("cwe: %d\n", akvs_jam2cws.FindCWE(flag.Args()[0], flag.Args()[1], bugs))
}
