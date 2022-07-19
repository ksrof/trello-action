/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package main

import "log"

const (
	name    string = "trello-action"
	version string = "1.0.0"
	author  string = "ksrof"
	env     string = "dev"
	cmd     string = "delete"
)

func main() {
	log.Printf(
		"Name: %s\n Version: %s\n Author: %s\n Env: %s\n Command: %s\n",
		name,
		version,
		author,
		env,
		cmd,
	)
}
