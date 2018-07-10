package main

import "fmt"
import "shared"

type ExtnMeta struct {
	data int
}

func (m ExtnMeta) Init(p shared.BR_v1) {
	if id, ok := p.ID(); ok {
		fmt.Println(p.Name() + ", " + id)
	} else {
		fmt.Println(p.Name() + ", access to ID not allowed")
	}
}

// exported
var PluginSymbol ExtnMeta
var Version string = shared.Version
