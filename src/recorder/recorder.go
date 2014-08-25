package main

import (
		"fmt"
		"os/exec"
		"strconv"
		"time"
		"flag"
)

var (
	recordTIme	= flag.Int("r", 30, "how long to record between iterations")
)

func main() {
	flag.Parse()
	for {
	for i :=0; i <10; i++ {
		cmd := exec.Command("/home/danl/Desktop/capture.sh", strconv.Itoa(i))
		go cmd.Start()
		time.Sleep(time.Duration(*recordTIme) * time.Second)
		cmd.Process.Kill()
		err := exec.Command("/bin/bash", "-c", "killall vlc").Run()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
		exec.Command("/bin/bash", "-c", "killall vlc").Run()
		time.Sleep(1 * time.Second)
	}
	}
}
