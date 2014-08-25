package main

import (
		"fmt"
		"os/exec"
		"strconv"
		"time"
		"flag"
)

var (
	recordTIme	= flag.Int("r", 1800, "how many seconds to record for each recording")
	numOfRecordings = flag.Int("n", 48, "Number of recording to keep before rolling over the first") 
	shellScript 	= flag.String("s", "", "shell script used to execute capture of stream.  Must take exactly one int argument")
	debug		= flag.Bool("v", false, "enable verbose output for debugging")
	
	TOOL_NAME = "FosRecorder"
	
)

// returns usage for phd_log_collector
func show_usage() {
	fmt.Printf("\nUsage for %s:\n\n", TOOL_NAME)

	// the printDefaults functions is ugly so changed the format here
	flag.VisitAll(func(flag *flag.Flag) {
		if flag.DefValue == "false" {
			fmt.Printf("  -%s\n     #~ %s\n\n", flag.Name, flag.Usage)
		} else {
			fmt.Printf("  -%s DEFAULT=\"%s\":\n     #~ %s\n\n", flag.Name, flag.DefValue, flag.Usage)
		}
	})
}

// killall vlc and sleep for given number of milliseconds
func killVlc(sec int) {
	err := exec.Command("/bin/bash", "-c", "killall vlc").Run()
	if err != nil && *debug {
		fmt.Println("Error when trying to kill vlc: ", err)
	}
	exec.Command("/bin/bash", "-c", "killall vlc").Run() // kill again just in case
	
	// make sure things have time to clean up before launch VLC again
	time.Sleep(time.Duration(sec) * time.Millisecond) 
}

func main() {
	flag.Usage = show_usage
	flag.Parse()
	for {
	for i :=0; i <*numOfRecordings; i++ {
		cmd := exec.Command(*shellScript, strconv.Itoa(i))
		go cmd.Start()
		time.Sleep(time.Duration(*recordTIme) * time.Second)
		cmd.Process.Kill()
		killVlc(500)
	}
	}
}
