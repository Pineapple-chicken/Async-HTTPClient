package main

import (
	"flag"
	"entities"
)

var n, o bool

func main()  {
	flag.BoolVar(&n, "n", false, "Using naive approach")
	flag.BoolVar(&o, "o", false, "Using optimized approach")
	flag.Parse()
	if n {
		synchronous.Naive()
	}
	if o {
		asynchronous.Optimized()
	}
}
