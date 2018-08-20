// This software is public domain. No rights are reserved. See LICENSE for more information.
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func octetGlobs(start byte, bits int) []string {
	min := int(start)
	max := int(start + byte((1<<uint(8-bits))-1))
	rv := []string{}

	// round down to a multiple of 10
	loopstart := (min / 10) * 10
	loopstop := max

	if min == 0 && max > 99 {
		rv = append(rv, "?")
		rv = append(rv, "??")
		loopstart = 100
	}

	if min <= 200 && max == 255 {
		loopstop = 200
	}

	for i := loopstart; i < loopstop; i += 10 {
		if i >= min && (i+9) <= max {
			if i == 0 {
				rv = append(rv, "?")
			} else {
				rv = append(rv, fmt.Sprintf("%d?", (i/10)))
			}
		} else if min < (i + 9) {
			var j int
			if min > i {
				j = min
			} else {
				j = i
			}
			for ; j <= max && j <= (i+9); j += 1 {
				rv = append(rv, fmt.Sprintf("%d", j))
			}
		}
	}

	if min <= 200 && max == 255 && loopstop == 200 {
		rv = append(rv, "2??")
	}

	return rv
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <cidr>\n", os.Args[0])
		os.Exit(127)
	}

	str := os.Args[1]
	ip, cidr, err := net.ParseCIDR(str)
	if err != nil {
		var ipErr error
		str += "/32"
		ip, cidr, ipErr = net.ParseCIDR(str)
		if ipErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
	}

	ip = ip.Mask(cidr.Mask)
	mask, _ := cidr.Mask.Size()

	var out string

	switch {
	case mask == 0:
		out = "*"

	case mask < 8:
		globs := octetGlobs(ip[0], mask)
		out = strings.Join(globs, ".* ") + ".*"

	case mask == 8:
		out = fmt.Sprintf("%d.*", ip[0])

	case mask < 16:
		globs := octetGlobs(ip[1], (mask - 8))
		out = fmt.Sprintf("%d.", ip[0]) + strings.Join(globs, fmt.Sprintf(".* %d.", ip[0])) + ".*"

	case mask == 16:
		out = fmt.Sprintf("%d.%d.*", ip[0], ip[1])

	case mask < 24:
		globs := octetGlobs(ip[2], (mask - 16))
		out = fmt.Sprintf("%d.%d.", ip[0], ip[1]) + strings.Join(globs, fmt.Sprintf(".* %d.%d.", ip[0], ip[1])) + ".*"

	case mask == 24:
		out = fmt.Sprintf("%d.%d.%d.*", ip[0], ip[1], ip[2])

	case mask < 32:
		globs := octetGlobs(ip[3], (mask - 24))
		out = fmt.Sprintf("%d.%d.%d.", ip[0], ip[1], ip[2]) + strings.Join(globs, fmt.Sprintf(".* %d.%d.%d.", ip[0], ip[1], ip[2])) + ".*"

	case mask == 32:
		out = ip.String()
	}

	fmt.Println(out)
}
