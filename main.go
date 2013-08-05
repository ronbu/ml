package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func MoveLink(from, to string) (err error) {
	from = filepath.Clean(from)
	to = filepath.Clean(to)
	absFrom, err := filepath.Abs(from)
	if err != nil {
		return
	}
	absTo, err := filepath.Abs(to)
	if err != nil {
		return
	}
	relTo, err := filepath.Rel(filepath.Dir(absFrom), absTo)
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
	err = os.Symlink(relTo, from)
	if err != nil {
		os.Rename(to, from)
	}
	return
}

func Reverse(link string) (target string, err error) {
	target, _ = os.Readlink(link)
	if target == "" {
		return
	}
	target = filepath.Join(filepath.Dir(link), target)
	err = os.Remove(link)
	if err != nil {
		return
	}
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

	if !(flag.NArg() == 2 || flag.NArg() == 1) {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Wrong number of arguments.")
		os.Exit(1)
	}

	var err error
	args := flag.Args()
	if *reverse {
		var target string
		target, err = Reverse(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		if target == "" {
			fmt.Fprintln(os.Stderr, args[0], ": is not a Symlink!")
		} else {
			fmt.Println(target)
		}
	} else {
		err = MoveLink(args[0], args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
}
