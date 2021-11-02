package main

import (
	"fmt"

	"github.com/discoriver/massh"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Create pointers to config & job
	config := massh.NewConfig()

	job := &massh.Job{
		Command: "echo hello world",
	}

	config.SetHosts([]string{"5.9.122.105"})

	// Password auth
	config.SetPasswordAuth("root", "Dell1074!")

	// Key auth in same config. Auth will try all methods provided before failing.
	err := config.SetPrivateKeyAuth("~/.ssh/id_rsa", "")
	if err != nil {
		panic(err)
	}

	config.SetJob(job)
	config.SetWorkerPool(2)
	config.SetSSHHostKeyCallback(ssh.InsecureIgnoreHostKey())

	// Make sure config will run
	config.CheckSanity()

	res, err := config.Run()
	if err != nil {
		panic(err)
	}

	for i := range res {
		fmt.Printf("%s:\n \t OUT: %s \t ERR: %s\n", res[i].Host, res[i].Output, res[i].Error)
	}
}
