package ogwc

import "regexp"

var crRegex = regexp.MustCompile("cr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var rrRegex = regexp.MustCompile("rr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")

var reportRegex = regexp.MustCompile("(rr|cr)-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var uidRegex = regexp.MustCompile("[0-f]{8}-([0-f]{4}-){3}[0-f]{12}")
