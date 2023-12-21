package kname

import (
	"fmt"
	"strings"
)

// Kname is a service that provides names tools
type Kname struct {
}

// NewKname creates a new Kname
func NewKname() *Kname {
	return &Kname{}
}

// GetFirstName returns the first name
func (k *Kname) GetShorten(name string, attempt int) string {
	name = strings.Trim(name, " ")
	if name == "" || attempt <= 0 {
		return ""
	}
	sname := strings.Split(strings.ToLower(name), " ")
	len := len(sname)
	pos := len - attempt
	if pos > 0 {
		return sname[0] + "_" + sname[pos]
	}
	rname := sname[0]
	if len > 1 {
		rname += "_" + sname[len-1]
		pos--
	}
	if pos < 0 {
		rname = rname + "_" + fmt.Sprintf("%d", pos*-1)
	}
	return rname
}
