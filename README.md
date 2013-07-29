[![Build Status](https://travis-ci.org/ronbu/ml.png?branch=master)](https://travis-ci.org/ronbu/ml)

# Installation

```bash
go get bitbucket.org/ronb/ml
```

ml is a very simple tool that combines *ln -s* and *mv*:

```bash
$ touch from
$ ml from to
$ ls -la
total 8
drwxr-xr-x  2 r  staff  136 Jul 29 20:53 .
drwxr-xr-x  4 r  staff  306 Jul 29 20:52 ..
lrwxr-xr-x  1 r  staff    2 Jul 29 20:53 from -> to
-rw-r--r--  1 r  staff    0 Jul 29 20:53 to
$ ml -reverse from
to
```