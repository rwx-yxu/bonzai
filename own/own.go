// Copyright 2022 bonzai-example Authors
// SPDX-License-Identifier: Apache-2.0

package own

import (
	"fmt"
	"log"

	Z "github.com/rwxrob/bonzai/z"
)

// private leaf
var Cmd = &Z.Cmd{
	Name:    `own`,
	Summary: "This is an own command",

	Call: func(caller *Z.Cmd, none ...string) error {
		val, _ := caller.Caller.C("some")
		fmt.Printf("%s\n", val)
		log.Print("I'm in my own file.")
		return nil
	},
}
