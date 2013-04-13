timealign
=========

timealign is a set of utility functions to perform alignments on Time structures.

Example
-------

	package main

	import (
	        "time"
	        "fmt"
	        "github.com/tadasv/timealign"
	)

	func main() {
	        t := time.Now()
	        fmt.Printf("Current time: %s\n", t)
	        fmt.Printf("1min alignment:  %s\n", timealign.AlignToMinute(t))
	        fmt.Printf("1min alignment:  %s\n", timealign.AlignToMinutes(t, 1))
	        fmt.Printf("2min alignment:  %s\n", timealign.AlignToMinutes(t, 2))
	        fmt.Printf("5min alignment:  %s\n", timealign.AlignToMinutes(t, 5))
	        fmt.Printf("hour alignemnt:  %s\n", timealign.AlignToHour(t))
	        fmt.Printf("day alignment:   %s\n", timealign.AlignToDay(t))
	        fmt.Printf("month alignment: %s\n", timealign.AlignToMonth(t))
	}

Output:

	Current time: 2013-04-13 16:49:59.244177 -0400 EDT
	1min alignment:  2013-04-13 16:49:00 -0400 EDT
	1min alignment:  2013-04-13 16:49:00 -0400 EDT
	2min alignment:  2013-04-13 16:48:00 -0400 EDT
	5min alignment:  2013-04-13 16:45:00 -0400 EDT
	hour alignemnt:  2013-04-13 16:00:00 -0400 EDT
	day alignment:   2013-04-13 00:00:00 -0400 EDT
	month alignment: 2013-03-31 00:00:00 -0400 EDT
