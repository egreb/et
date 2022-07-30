package main

import (
	"fmt"
	"os"
	"strconv"
)

func debug(values ...interface{}) {
	fmt.Printf("%#v\n", values)
}

func printErr(err error) {
	fmt.Println(fmt.Errorf("%w", err))
}

func printWithTab(v string, level int) {
	output := ""
	for i := 0; i < level; i++ {
		output += "\t"
	}

	fmt.Printf("%s%s\n", output, v)
}

func printPath(path string, level int, maxLevel int) {
	if maxLevel != -1 && level >= maxLevel {
		return
	}

	ff, err := os.ReadDir(path)
	if err != nil {
		printErr(err)
		return
	}

	for _, f := range ff {
		n := f.Name()
		printWithTab(n, level)
		if !f.IsDir() {
			continue
		}

		printPath(fmt.Sprintf("%s/%s", path, n), level+1, maxLevel)
	}
}

func main() {
	args := os.Args[1:]
	dir := ""
	if len(args) == 0 {
		p, err := os.Getwd()
		if err != nil {
			printErr(err)
			return
		}

		dir = p
	} else {

		dir = args[0]
	}

	maxLevel := -1
	if len(args) == 2 {
		l, err := strconv.Atoi(args[1])
		if err != nil {
			printErr(err)
			return
		}
		maxLevel = l
	}

	printPath(dir, 0, maxLevel)
}
