package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func MoveLink(from, to string) (err error) {
	relTo, err := filepath.Rel(filepath.Dir(from), to)
	if err != nil {
		return
	}
	target, _ := os.Readlink(from)
	if target != "" && relTo == target {
		return
	}
	err = os.Rename(from, to)
	if err != nil {
		return
	}
	// println("relto target: ", relTo, target)
	// println("from to: ", from, to)
	return os.Symlink(relTo, from)
}

func Reverse(link string) (target string, err error) {
	target, _ = os.Readlink(link)
	if target == "" {
		return
	}
	target = filepath.Join(filepath.Dir(link), target)
	err = os.Rename(target, link)
	return
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [-r] from [to]:\n", os.Args[0])
		flag.PrintDefaults()
	}
	reverse := flag.Bool("reverse", false, "Reverse ml")
	flag.Parse()

	var err error
	args := flag.Args()
	if *reverse {
		if flag.NArg() != 1 {
			flag.Usage()
			fmt.Fprintln(os.Stderr, "Needs 2 arguments when -reverse is not set.")
			os.Exit(1)
		}
		var target string
		target, err = Reverse(args[0])
		if target == "" {
			fmt.Fprintln(os.Stderr, args[0], ": is not a Symlink!")
		} else {
			fmt.Println(target)
		}
	} else {
		if flag.NArg() != 2 {
			flag.Usage()
			fmt.Fprintln(os.Stderr, "Only 1 argument when -reverse is set.")
			os.Exit(1)
		}
		err = MoveLink(args[0], args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
