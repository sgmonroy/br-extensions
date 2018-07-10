package main

import (
	"fmt"
	"os"
	"plugin"
	"shared"
)

type Base struct {
	name string
}

func (b Base) Name() string {
	return b.name
}

func (b Base) ID() (string, bool) {
	return "", false
}

type Extended struct {
	Base
	id string
}

func (e Extended) Name() string {
	return e.name
}

func (e Extended) ID() (string, bool) {
	return e.id, true
}

func main() {
	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open("./plugins/plugin.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symVersion, err := plug.Lookup("Version")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ver, ok := symVersion.(*string)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		fmt.Printf("%T", symVersion)
		os.Exit(1)
	}
	fmt.Println("PluginVersion: " + *ver)

	// 3. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symPlugin, err := plug.Lookup("PluginSymbol")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 4. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	p, ok := symPlugin.(shared.Plugin_v1)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		fmt.Printf("%T", symVersion)
		os.Exit(1)
	}

	data := Extended{}
	data.name = "JohnDoe"
	data.id = "unknown"
	// 5. use the module with base access
	p.Init(data.Base)

	// 6. use the module with extended access
	p.Init(data)
}
