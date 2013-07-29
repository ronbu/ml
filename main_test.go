package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveLinkAndReverse(t *testing.T) {
	tmp, rm := TempDir()
	defer rm()

	from := tmp + "/from"
	to := tmp + "/to"
	err := ioutil.WriteFile(from, []byte(""), 0777)
	if err != nil {
		t.Fatal(err)
	}
	err = MoveLink(from, to)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(to)
	if err != nil {
		t.Fatal(err)
	}
	link, err := os.Readlink(from)
	if err != nil {
		t.Fatal(err)
	}
	if link != "to" {
		t.Fatal("Wrong link target created: ", link)
	}

	target, err := Reverse(from)
	if err != nil {
		t.Fatal(err)
	}
	if target != to {
		t.Fatal("Reverse reports wrong link target: ", target)
	}
}

func TestExistingLink(t *testing.T) {
	tmp, rm := TempDir()
	defer rm()

	from := tmp + "/from"
	to := tmp + "/to"
	err := ioutil.WriteFile(to, []byte(""), 0777)
	if err != nil {
		t.Fatal(err)
	}
	os.Symlink("to", from)
	if err != nil {
		t.Fatal(err)
	}

	err = MoveLink(from, to)
	if err != nil {
		t.Fatal(err)
	}

	// MoveLink should do nothing if {from}
	// is already a symlink to {to}
	link, err := os.Readlink(from)
	if err != nil {
		t.Fatal(err)
	}
	if link != "to" {
		t.Fatal("MoveLink has changed existing link to: ", link)
	}
	_, err = os.Stat(to)
	if err != nil {
		t.Fatal("MoveLink deleted file at: ", to)
	}
}

func TestReverseOnNonLink(t *testing.T) {
	tmp, rm := TempDir()
	defer rm()

	fp := tmp + "/file"
	err := ioutil.WriteFile(fp, []byte(""), 0777)
	target, err := Reverse(fp)
	if err != nil {
		t.Fatal(err)
	}
	if target != "" {
		t.Fatal("{target} should be empty when file is not a symlink: ", target)
	}
}

func TempDir() (string, func()) {
	path, err := ioutil.TempDir("", "TestMl")
	check(err)
	path, err = filepath.EvalSymlinks(path)
	check(err)
	return path, func() {
		check(os.RemoveAll(path))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
