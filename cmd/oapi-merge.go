package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/badygin-yakov/oapi-merge/pkg/merge"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec files")
		os.Exit(1)
	}

	filePaths := strings.Split(flag.Arg(0), ",")

	err := merge.Do(filePaths, flag.Arg(1))
	if err != nil {
		log.Fatalf("failed merge files. %v", err)
	}
}
