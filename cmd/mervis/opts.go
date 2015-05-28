package main

import "flag"

type opts struct {
	Addr       string
	ClientRoot string
}

func getopts() opts {
	root := flag.String("r", "", "the root path of the client content. blank will use bindata")
	flag.Parse()

	return opts{
		Addr:       ":3001",
		ClientRoot: *root,
	}
}
